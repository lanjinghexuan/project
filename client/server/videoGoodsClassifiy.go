package server

import (
	"context"
	"github.com/lanjinghexuan/project/common/model"
	pb "github.com/lanjinghexuan/project/common/proto/videoGoodsClassifiy"
)

type GoodsClassifiyServer struct {
	pb.UnimplementedGoodsClassifiyServer
}

func (g GoodsClassifiyServer) GetGoodsClass(ctx context.Context, req *pb.GoodsClassReq) (*pb.GoodsClassRes, error) {
	var class model.VideoGoodsClassifiy
	var err error
	list, err := class.GetClass(req.GetPid())
	if err != nil {
		return nil, err
	}
	var resp []*pb.GoodsClass
	for _, v := range list {
		resp = append(resp, &pb.GoodsClass{
			Id:            v.Id,
			ClassifiyName: v.ClassifiyName,
			Pid:           v.Pid,
			Soft:          v.Sort,
		})
	}
	return &pb.GoodsClassRes{Goodsclass: resp}, err

}
