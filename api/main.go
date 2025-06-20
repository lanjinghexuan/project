package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lanjinghexuan/project/common/init"
)

func main() {
	//add := map[string]string{
	//	"Id":   "4",
	//	"Name": "测试数据赵六",
	//}
	//pkr.AddEs(add)
	//pkr.SearchEs()
	//pkr.DelEs()
	r := gin.Default()
	//router.Router(r)

	//调用文心一言代码
	//r.GET("/sendgpt", handle.SendGpt)
	//r.GET("/getgptdata", handle.GetGptData)
	//r.GET("/sendFlowGpt", handle.SendFlowGpt)
	//r.GET("/getFlowGpt", handle.GetFlowGpt)

	r.Run(":8080")
}
