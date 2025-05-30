
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ProductSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
    request.PageInfo
}
