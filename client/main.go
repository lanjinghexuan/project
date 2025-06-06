package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"project/client/server"
	"project/common/gload"
	_ "project/common/init"
	pb "project/common/proto/videoGoodsClassifiy"
	pb2 "project/common/proto/videoUser"
)

func main() {
	flag.Parse()
	fmt.Println(fmt.Sprintf("%s:%d", gload.CONFIG.Server.Host, gload.CONFIG.Server.Port))
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", gload.CONFIG.Server.Host, gload.CONFIG.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGoodsClassifiyServer(s, server.GoodsClassifiyServer{})
	pb2.RegisterVideoUserServer(s, server.VideoUserServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
