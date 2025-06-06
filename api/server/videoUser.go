package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"project/common/gload"
	pb "project/common/proto/videoUser"
)

type VideoUser func(ctx context.Context, in pb.VideoUserClient) (interface{}, error)

func Client(ctx context.Context, clients VideoUser) (interface{}, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", gload.CONFIG.Server.Host, gload.CONFIG.Server.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewVideoUserClient(conn)
	return clients(ctx, c)
}

func Login(ctx context.Context, client *pb.LoginRep) (*pb.LoginRes, error) {
	res, err := Client(ctx, func(ctx context.Context, in pb.VideoUserClient) (interface{}, error) {
		login, err := in.Login(ctx, client)
		if err != nil {
		}
		return login, err
	})
	if err != nil {
	}
	return res.(*pb.LoginRes), err
}
