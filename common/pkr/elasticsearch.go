package pkr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lanjinghexuan/project/common/gload"
	"strings"
)

// CreateEs 创建 Elasticsearch 索引
func CreateEs(index string) (bool, error) {
	resp, err := gload.ES.Indices.Create(index)
	if err != nil {
		return false, fmt.Errorf("create index failed, err: %v", err)
	}
	if resp.IsError() {
		return false, fmt.Errorf("create index failed, err: %v", resp)
	}
	return true, nil
}

type AddVideo struct {
	Id    int64
	Title string
}

// AddEs 向 Elasticsearch 添加文档
func AddEs(index string, add interface{}, id string) (bool, error) {
	addjson, err := json.Marshal(add)
	if err != nil {
		return false, fmt.Errorf("error marshaling document: %v", err)
	}
	res, err := gload.ES.Index(index, strings.NewReader(string(addjson)), gload.ES.Index.WithDocumentID(id))
	if err != nil {
		return false, err
	}
	if res.IsError() {
		return false, fmt.Errorf("add document failed, err: %v", res)
	}
	defer res.Body.Close()
	return true, nil
}

// BulkAddEs 批量添加文档到 Elasticsearch
func BulkAddEs(index string, documents []interface{}) (bool, error) {
	var bulkRequest bytes.Buffer

	for _, doc := range documents {
		bulkRequest.WriteString(fmt.Sprintf(`{"index": {"_index": "%s"}}`, index))
		bulkRequest.WriteString("\n")
		jsonData, err := json.Marshal(doc)
		if err != nil {
			return false, fmt.Errorf("error marshaling document: %v", err)
		}
		bulkRequest.Write(jsonData)
		bulkRequest.WriteString("\n")
	}

	res, err := gload.ES.Bulk(strings.NewReader(bulkRequest.String()))
	if err != nil {
		return false, fmt.Errorf("error sending bulk request: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var response map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			return false, fmt.Errorf("error decoding response: %v", err)
		}
		return false, fmt.Errorf("error in bulk request: %v", response["error"])
	}

	return true, nil
}

// SearchEs 在 Elasticsearch 中搜索文档
func SearchEs(query string, indexName string) ([]interface{}, error) {
	var r map[string]interface{}
	res, err := gload.ES.Search(
		gload.ES.Search.WithIndex(indexName),
		gload.ES.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return nil, fmt.Errorf("search index failed, err: %v", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("decode failed, err: %v", err)
	}

	var newResult []interface{}
	if r["hits"] == nil {
		return nil, nil
	}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		newResult = append(newResult, hit.(map[string]interface{})["_source"])
	}

	return newResult, nil
}

// DelEs 从 Elasticsearch 删除文档
func DelEs() error {
	query := `{"query": {
        "match": {
            "Id": "3"
        }
    }}`
	res, err := gload.ES.DeleteByQuery([]string{"video_works"}, strings.NewReader(query))
	if err != nil {
		return fmt.Errorf("elasticsearch error: %v", err)
	}
	if res.StatusCode == 200 {
		fmt.Println("删除成功.success")
	} else {
		fmt.Printf("删除失败: %v\n", res)
	}
	return nil
}

// IssetIndex 检查 Elasticsearch 索引是否存在
func IssetIndex(index string) (bool, error) {
	res, err := gload.ES.Indices.Exists([]string{index}, gload.ES.Indices.Exists.WithContext(gload.Ctx))
	if err != nil {
		return false, fmt.Errorf("IssetIndex 获取错误。error: %v", err)
	}
	if res.IsError() {
		return false, nil
	}
	return true, nil
}

// DelIndex 删除 Elasticsearch 索引
func DelIndex(index string) (bool, error) {
	res, err := gload.ES.Indices.Delete([]string{index})
	if err != nil {
		return false, err
	}
	fmt.Println("DelIndex, res:", res)
	return true, nil
}
