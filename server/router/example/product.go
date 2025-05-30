package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProductRouter struct {}

// InitProductRouter 初始化 product表 路由信息
func (s *ProductRouter) InitProductRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	productRouter := Router.Group("product").Use(middleware.OperationRecord())
	productRouterWithoutRecord := Router.Group("product")
	productRouterWithoutAuth := PublicRouter.Group("product")
	{
		productRouter.POST("createProduct", productApi.CreateProduct)   // 新建product表
		productRouter.DELETE("deleteProduct", productApi.DeleteProduct) // 删除product表
		productRouter.DELETE("deleteProductByIds", productApi.DeleteProductByIds) // 批量删除product表
		productRouter.PUT("updateProduct", productApi.UpdateProduct)    // 更新product表
	}
	{
		productRouterWithoutRecord.GET("findProduct", productApi.FindProduct)        // 根据ID获取product表
		productRouterWithoutRecord.GET("getProductList", productApi.GetProductList)  // 获取product表列表
	}
	{
	    productRouterWithoutAuth.GET("getProductPublic", productApi.GetProductPublic)  // product表开放接口
	}
}
