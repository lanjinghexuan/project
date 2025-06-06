package main

import (
	"fmt"
	_ "project/common/cmd"
	"project/common/gload"
	"project/common/pkr"
)

type videoWorks struct {
	Id    int32
	Title string
}

/*
同步mysql到es的示例
*/
func main() {
	//pkr.DelIndex("video_works")
	//return

	//判断索引是否存在
	res, err := pkr.IssetIndex("video_works")
	if err != nil {
		fmt.Println(err)
	}
	if !res {
		fmt.Println("索引不存在，创建索引")
		res, err = pkr.CreateEs("video_works")
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
	last, err := pkr.SearchEs(query)
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
	var videoworks []*videoWorks
	err = gload.DB.Table("video_works").Where("id > ?", id).Find(&videoworks).Error
	if err != nil {
		fmt.Println("===================mysql error:", err)
		return
	}
	fmt.Println(videoworks)

	//BulkAddEs(Ai编写的批量添加，在大数据使用)
	var bulkRequest []interface{}
	for _, doc := range videoworks {
		bulkRequest = append(bulkRequest, doc)
	}
	res, err = pkr.BulkAddEs("video_works", bulkRequest)
	fmt.Println(res, err)
	return

	//添加es(没有做批量添加使用的是循环添加单条)
	for _, v := range videoworks {
		if v.Id == 0 {
			fmt.Println("数据添加完毕")
			return
		}
		r, err := pkr.AddEs("video_works", v)
		if err != nil {
			fmt.Println("addes error:", err)
			return
		}
		if !r {
			fmt.Println("addes error:", err)
			return
		}
	}

	fmt.Println("添加成功")
}
