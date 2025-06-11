package main

import (
	"fmt"
	_ "project/common/cmd"
	"project/common/gload"
)

type videoWorks struct {
	Id    int32
	Title string
}

func main() {
	coll := gload.MONGO.Database("db").Collection("test")

	// 操作mongo进行添加
	var videoworks []*videoWorks
	err := gload.DB.Table("video_works").Where("id > ?", 5).Find(&videoworks).Error
	if err != nil {
		fmt.Println(err)
	}
	var docs []interface{}
	for _, v := range videoworks {
		docs = append(docs, v)
	}

	_, err = coll.InsertMany(gload.Ctx, docs)
	if err != nil {
		fmt.Println(err)
	}

	//操作mongo查询
	//options1 := options.Find()
	////排序
	//options1.Sort = bson.D{{"id", -1}}
	////开始位置
	//options1.SetSkip(2)
	////截取个数
	//options1.SetLimit(5)
	////条件(建议借用ai)
	//fillter := bson.M{"id": bson.M{"$gt": 5}}
	//
	//res, err := coll.Find(gload.Ctx, fillter, options1)
	//if err != nil {
	//	fmt.Println("error:", err)
	//	return
	//}
	//var results []videoWorks
	//res.All(gload.Ctx, &results)
	//for _, result := range results {
	//	//展示数据
	//	fmt.Printf("%+v\n", result)
	//}

	//操作mongo修改
	//update := bson.M{"$set": bson.M{
	//	"title": "阳燧成辉",
	//}}
	//
	//res, err := coll.UpdateMany(gload.Ctx, bson.M{"id": bson.M{"$lte": 6}}, update)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//if res.MatchedCount > 0 {
	//	fmt.Printf("updated %v documents\n", res.ModifiedCount)
	//}

	//操作mongo删除
	//where := bson.M{"title": "第七个测试"}
	//res, err := coll.DeleteMany(gload.Ctx, where)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//if res.DeletedCount > 0 {
	//	fmt.Println("Delete 成功")
	//}
}
