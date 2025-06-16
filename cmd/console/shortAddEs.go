package console

import (
	"fmt"
	"project/common/gload"
	"project/common/pkr"
	"time"
)

type Short struct {
	Id        string
	Name      string
	TypeId    int32
	CreatedAt time.Time
}

func AddEs() {
	indexName := "short_video"
	tableName := "short_video"
	res, err := pkr.IssetIndex(indexName)
	if err != nil {
		fmt.Println(err)
	}
	if !res {
		fmt.Println("索引不存在，创建索引")
		res, err = pkr.CreateEs(indexName)
		if err != nil {
			fmt.Println("创建索引失败.error:", err)
			return
		}
		if !res {
			fmt.Println("创建索引失败.error:", err)
			return
		}

	}

	//查询最后一条插入的数据
	query := `{
		"sort": [
		  {
			"Id": {
			  "order": "desc"
			}
		  }
		],
		"size": 1
	}`
	last, err := pkr.SearchEs(query, indexName)
	if err != nil {
		fmt.Println(err)
		return
	}
	var id int32 = 0
	if last != nil {
		for _, item := range last {
			newItem := item.(map[string]interface{})
			id = int32(newItem["Id"].(float64))
		}
	}

	//获取到最后一条插入的数据查询数据库
	//如果数据量过大必须分页处理
	var datas []*Short
	err = gload.DB.Table(tableName).Where("id > ?", id).Find(&datas).Error
	if err != nil {
		fmt.Println("===================mysql error:", err)
		return
	}
	fmt.Println(datas)
	if len(datas) == 0 {
		fmt.Println("数据添加完毕")
		return
	}

	//BulkAddEs(Ai编写的批量添加，在大数据使用)
	var bulkRequest []interface{}
	for _, doc := range datas {
		bulkRequest = append(bulkRequest, doc)
	}
	res, err = pkr.BulkAddEs(indexName, bulkRequest)
	fmt.Println(res, err)
	return
}
