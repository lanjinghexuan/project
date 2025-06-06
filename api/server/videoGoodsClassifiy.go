package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"project/common/gload"
	pb "project/common/proto/videoGoodsClassifiy"
)

type videoGoodsClassCHand func(ctx context.Context, client pb.GoodsClassifiyClient) (interface{}, error)

func videoGoodsClassifiyClient(ctx context.Context, client videoGoodsClassCHand) (interface{}, error) {
	coon, err := grpc.Dial(fmt.Sprintf("%s:%d", gload.CONFIG.Server.Host, gload.CONFIG.Server.Port), grpc.WithInsecure())
	if err != nil {
	}
	defer coon.Close()
	c := pb.NewGoodsClassifiyClient(coon)
	return client(ctx, c)
}

func GetGoodsClass(ctx context.Context, in *pb.GoodsClassReq) (*pb.GoodsClassRes, error) {
	res, err := videoGoodsClassifiyClient(ctx, func(ctx context.Context, client pb.GoodsClassifiyClient) (interface{}, error) {
		resp, err := client.GetGoodsClass(ctx, in)
		return resp, err
	})
	if err != nil {
	}
	return res.(*pb.GoodsClassRes), nil
}
