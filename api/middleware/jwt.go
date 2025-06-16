package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lanjinghexuan/project/common/pkr"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(200, gin.H{
				"code":    601,
				"message": "token is empty",
			})
		}
		data, err := pkr.ParseJwt(token)
		if err != nil {
			c.JSON(200, gin.H{
				"code":    602,
				"message": "token is invalid",
			})
		}
		fmt.Println(data)
		userid := data["Id"]
		fmt.Println(data, userid)
		c.Set("userid", userid)
		// 请求前

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}

//	router.Use(Logger())
