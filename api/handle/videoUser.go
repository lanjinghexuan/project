package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project/api/request"
	"project/api/server"
	"project/common/pkr"
	pb "project/common/proto/videoUser"
)

func Login(c *gin.Context) {
	var req request.LoginReq
	err := c.ShouldBind(&req)
	if err != nil {
		fmt.Println("参数接收失败 .error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 600,
			"msg":  err,
		})
		return
	}
	loginreq := &pb.LoginRep{
		Name:     req.Username,
		UserCode: req.Password,
	}
	data, err := server.Login(c, loginreq)
	if err != nil {
		fmt.Println("调用服务层读取信息失败.error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 600,
			"msg":  err,
		})
		return
	}
	token, err := pkr.GenerateJWT(data.Id)
	if err != nil {
		fmt.Println("jwt生成失败.error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 600,
			"msg":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": token,
	})
	return
}
