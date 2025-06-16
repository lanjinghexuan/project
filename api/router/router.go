package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lanjinghexuan/project/api/handle"
	"github.com/lanjinghexuan/project/api/middleware"
)

func Router(r *gin.Engine) {
	c := r.Group("api")
	{
		c.GET("login", handle.Login)
		c.Use(middleware.Logger())
		c.GET("goodsclass", func(c *gin.Context) {})
	}
}
