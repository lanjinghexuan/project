package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lanjinghexuan/project/api/request"
	"github.com/lanjinghexuan/project/api/server"
	pb "github.com/lanjinghexuan/project/common/proto/videoGoodsClassifiy"
)

func GetGoodsClass(c *gin.Context) {
	var req request.GetGoodsClassifiyReq
	err := c.ShouldBind(&req)
	if err != nil {
	}

	res, err := server.GetGoodsClass(c, &pb.GoodsClassReq{
		Pid: req.Pid,
	})

	if err != nil {

	}
	fmt.Println(res)

}
