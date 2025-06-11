package pkr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"project/common/gload"
	"strings"
)

//没有封装修改，建议删除在添加

func CreateEs(index string) (bool, error) {
	resp, err := gload.ES.Indices.
		Create(index)
	if err != nil {
		fmt.Printf("create index failed, err:%v\n", err)
		return false, err
	}
	if resp.IsError() {
		fmt.Printf("create index failed, err:%v\n", resp)
	}
	return true, err
}

type AddVideo struct {
	Id    int64
	Title string
}

func AddEs(index string, add interface{}, id string) (bool, error) {
	addjson, _ := json.Marshal(add)
	res, err := gload.ES.Index(index, strings.NewReader(string(addjson)), gload.ES.Index.WithDocumentID(id))
	if err != nil {
		return false, err
	}
	if res.IsError() {
		return false, err
	}
	defer res.Body.Close()
	return true, nil
}

// BulkAddEs 批量添加文档到 Elasticsearch (Ai编写)
func BulkAddEs(index string, documents []interface{}) (bool, error) {
	var bulkRequest bytes.Buffer

	for _, doc := range documents {
		// 添加索引操作
		bulkRequest.WriteString(fmt.Sprintf(`{"index": {"_index": "%s"}}`, index))
		bulkRequest.WriteString("\n")
		// 添加文档数据
		jsonData, err := json.Marshal(doc)
		if err != nil {
			return false, fmt.Errorf("error marshaling document: %v", err)
		}
		bulkRequest.Write(jsonData)
		bulkRequest.WriteString("\n")
	}

	// 打印请求体以调试
	fmt.Println("Bulk request body:", bulkRequest.String())

	// 执行批量请求
	res, err := gload.ES.Bulk(strings.NewReader(bulkRequest.String()))
	if err != nil {
		return false, fmt.Errorf("error sending bulk request: %v", err)
	}
	defer res.Body.Close()

	// 检查响应是否有错误
	if res.IsError() {
		var response map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			return false, fmt.Errorf("error decoding response: %v", err)
		}
		return false, fmt.Errorf("error in bulk request: %v", response["error"])
	}

	return true, nil
}

func SearchEs(query string) ([]interface{}, error) {
	var r map[string]interface{}
	gload.ES.Search.WithIndex("video_works")
	res, err := gload.ES.Search(
		gload.ES.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		fmt.Printf("search index failed, err:%v\n", err)
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		fmt.Printf("decode failed, err:%v\n", err)
		return nil, err
	}
	var newResult []interface{}
	if r["hits"] == nil {
		return nil, nil
	}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		newResult = append(newResult, hit.(map[string]interface{})["_source"])
	}

	return newResult, err
}

func DelEs() {
	query := `{"query": {
    "match": {
      "Id": "3"
    }
  }}`
	res, err := gload.ES.DeleteByQuery([]string{"video_works"}, strings.NewReader(query))
	//示例
	//res, err = gload.ES.Delete("video_works", "5")
	if err != nil {
		fmt.Println("<UNK>elasticsearch<UNK>.error:", err)
	}
	if res.StatusCode == 200 {
		fmt.Println("删除成功.success")
	}
	fmt.Println(res, "删除失败")
}

func IssetIndex(index string) (bool, error) {
	res, err := gload.ES.Indices.Exists([]string{index}, gload.ES.Indices.Exists.WithContext(gload.Ctx))
	if err != nil {
		fmt.Println("============= IssetIndex获取错误。error:", res)
		return false, err
	}
	if res.IsError() {
		fmt.Println("============= IssetIndex获取错误。error:", res)
		return false, nil
	}
	return true, nil
}

func DelIndex(index string) (bool, error) {
	res, err := gload.ES.Indices.Delete([]string{index})
	if err != nil {
		return false, err
	}
	fmt.Println("DelIndex,res:", res)
	return true, nil
}
