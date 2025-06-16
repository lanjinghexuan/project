package server

import (
	"context"
	"github.com/lanjinghexuan/project/common/gload"
	pb "github.com/lanjinghexuan/project/common/proto/videoUser"
)

type VideoUserServer struct {
	pb.UnimplementedVideoUserServer
}

func (c VideoUserServer) Login(ctx context.Context, rep *pb.LoginRep) (res *pb.LoginRes, err error) {
	err = gload.DB.Table("video_user").Where("name = ?", rep.Name).Where("user_code = ?", rep.UserCode).Limit(1).First(&res).Error
	return res, err
}
