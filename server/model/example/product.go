
// 自动生成模板Product
package example
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// product表 结构体  Product
type Product struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"column:name;size:50;"`  //name字段
  Price  *float64 `json:"price" form:"price" gorm:"column:price;size:10;"`  //price字段
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName product表 Product自定义表名 product
func (Product) TableName() string {
    return "product"
}





