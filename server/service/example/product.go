
package example

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
    exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
    "gorm.io/gorm"
)

type ProductService struct {}
// CreateProduct 创建product表记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService) CreateProduct(ctx context.Context, product *example.Product) (err error) {
	err = global.GVA_DB.Create(product).Error
	return err
}

// DeleteProduct 删除product表记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService)DeleteProduct(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&example.Product{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&example.Product{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteProductByIds 批量删除product表记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService)DeleteProductByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&example.Product{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&example.Product{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateProduct 更新product表记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService)UpdateProduct(ctx context.Context, product example.Product) (err error) {
	err = global.GVA_DB.Model(&example.Product{}).Where("id = ?",product.ID).Updates(&product).Error
	return err
}

// GetProduct 根据ID获取product表记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService)GetProduct(ctx context.Context, ID string) (product example.Product, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&product).Error
	return
}
// GetProductInfoList 分页获取product表记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService)GetProductInfoList(ctx context.Context, info exampleReq.ProductSearch) (list []example.Product, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&example.Product{})
    var products []example.Product
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&products).Error
	return  products, total, err
}
func (productService *ProductService)GetProductPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
