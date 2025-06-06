package router

import (
	"github.com/gin-gonic/gin"
	"project/api/handle"
	"project/api/middleware"
)

func Router(r *gin.Engine) {
	c := r.Group("api")
	{
		c.GET("login", handle.Login)
		c.Use(middleware.Logger())
		c.GET("goodsclass", func(c *gin.Context) {})
	}
}
