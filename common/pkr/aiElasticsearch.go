package pkr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Config 配置结构
type Config struct {
	Addresses  []string
	Username   string
	Password   string
	MaxRetries int
	Timeout    time.Duration
}

// Client ES客户端封装
type Client struct {
	es      *elasticsearch.Client
	logger  *zap.Logger
	timeout time.Duration
}

// NewClient 创建ES客户端
func NewClient(cfg Config, logger *zap.Logger) (*Client, error) {
	cfgAddresses := viper.GetStringSlice("elasticsearch.addresses")
	if len(cfgAddresses) > 0 {
		cfg.Addresses = cfgAddresses
	}

	esCfg := elasticsearch.Config{
		Addresses:  cfg.Addresses,
		Username:   cfg.Username,
		Password:   cfg.Password,
		MaxRetries: cfg.MaxRetries,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: cfg.Timeout,
		},
	}

	es, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		return nil, fmt.Errorf("创建ES客户端失败: %w", err)
	}

	// 测试连接
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := es.Info()
	if err != nil {
		return nil, fmt.Errorf("ES连接测试失败: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("ES返回错误: %s", res.Status())
	}

	return &Client{
		es:      es,
		logger:  logger,
		timeout: cfg.Timeout,
	}, nil
}

// VideoWork 视频作品模型
type VideoWork struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// IndexService 索引服务
type IndexService struct {
	client *Client
}

// NewIndexService 创建索引服务
func NewIndexService(client *Client) *IndexService {
	return &IndexService{client: client}
}

// CreateIndex 创建索引
func (s *IndexService) CreateIndex(index string) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.client.timeout)
	defer cancel()

	req := esapi.IndicesCreateRequest{
		Index: index,
		Body: strings.NewReader(`{
			"mappings": {
				"properties": {
					"id":    { "type": "long" },
					"title": { "type": "text", "analyzer": "ik_max_word" }
				}
			}
		}`),
	}

	res, err := req.Do(ctx, s.client.es)
	if err != nil {
		return fmt.Errorf("创建索引失败: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return parseErrorResponse(res, "创建索引失败")
	}

	return nil
}

// DocumentService 文档服务
type DocumentService struct {
	client *Client
}

// NewDocumentService 创建文档服务
func NewDocumentService(client *Client) *DocumentService {
	return &DocumentService{client: client}
}

// AddDocument 添加文档
func (s *DocumentService) AddDocument(index string, doc interface{}, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.client.timeout)
	defer cancel()

	docJSON, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("序列化文档失败: %w", err)
	}

	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       bytes.NewReader(docJSON),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, s.client.es)
	if err != nil {
		return fmt.Errorf("添加文档失败: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return parseErrorResponse(res, "添加文档失败")
	}

	return nil
}

// BulkAddDocuments 批量添加文档
func (s *DocumentService) BulkAddDocuments(index string, docs []interface{}) error {
	if len(docs) == 0 {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.client.timeout*time.Duration(len(docs)/100+1))
	defer cancel()

	var bulkReq bytes.Buffer
	for _, doc := range docs {
		docJSON, err := json.Marshal(doc)
		if err != nil {
			return fmt.Errorf("序列化文档失败: %w", err)
		}

		meta := fmt.Sprintf(`{"index": {"_index": "%s"}}%s`, index, "\n")
		bulkReq.WriteString(meta)
		bulkReq.Write(docJSON)
		bulkReq.WriteString("\n")
	}

	req := esapi.BulkRequest{
		Body:    bytes.NewReader(bulkReq.Bytes()),
		Refresh: "true",
	}

	res, err := req.Do(ctx, s.client.es)
	if err != nil {
		return fmt.Errorf("批量添加文档失败: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return parseErrorResponse(res, "批量添加文档失败")
	}

	// 检查批量操作结果
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return fmt.Errorf("解析批量操作结果失败: %w", err)
	}

	if failures, ok := result["errors"].(bool); ok && failures {
		if items, ok := result["items"].([]interface{}); ok {
			for _, item := range items {
				if obj, ok := item.(map[string]interface{}); ok {
					for _, v := range obj {
						if errObj, ok := v.(map[string]interface{}); ok {
							if errObj["error"] != nil {
								s.client.logger.Warn("批量操作失败项",
									zap.Any("error", errObj["error"]))
							}
						}
					}
				}
			}
		}
		return fmt.Errorf("部分文档添加失败")
	}

	return nil
}

// SearchService 搜索服务
type SearchService struct {
	client *Client
}

// NewSearchService 创建搜索服务
func NewSearchService(client *Client) *SearchService {
	return &SearchService{client: client}
}

// Search 搜索文档
func (s *SearchService) Search(index, query string) ([]VideoWork, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.client.timeout)
	defer cancel()

	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  strings.NewReader(query),
	}

	res, err := req.Do(ctx, s.client.es)
	if err != nil {
		return nil, fmt.Errorf("搜索失败: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, parseErrorResponse(res, "搜索失败")
	}

	var result struct {
		Hits struct {
			Hits []struct {
				Source VideoWork `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析搜索结果失败: %w", err)
	}

	var docs []VideoWork
	for _, hit := range result.Hits.Hits {
		docs = append(docs, hit.Source)
	}

	return docs, nil
}

// DeleteService 删除服务
type DeleteService struct {
	client *Client
}

// NewDeleteService 创建删除服务
func NewDeleteService(client *Client) *DeleteService {
	return &DeleteService{client: client}
}

// DeleteByQuery 根据查询删除文档
func (s *DeleteService) DeleteByQuery(index, query string) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.client.timeout)
	defer cancel()

	req := esapi.DeleteByQueryRequest{
		Index: []string{index},
		Body:  strings.NewReader(query),
	}

	res, err := req.Do(ctx, s.client.es)
	if err != nil {
		return fmt.Errorf("删除失败: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return parseErrorResponse(res, "删除失败")
	}

	return nil
}

// 辅助函数：解析错误响应
func parseErrorResponse(res *esapi.Response, prefix string) error {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("%s: %s (读取响应体失败: %v)", prefix, res.Status(), err)
	}

	var errorResponse struct {
		Error struct {
			Reason string `json:"reason"`
		} `json:"error"`
	}

	if json.Unmarshal(body, &errorResponse) == nil && errorResponse.Error.Reason != "" {
		return fmt.Errorf("%s: %s (原因: %s)", prefix, res.Status(), errorResponse.Error.Reason)
	}

	return fmt.Errorf("%s: %s (响应: %s)", prefix, res.Status(), string(body))
}
