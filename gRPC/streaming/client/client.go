package main

import (
	"context"
	"net"

	"google.golang.org/grpc"

	pb "go-third-party/gRPC/streaming/stream"
	"log"
)

const (
	PORT = ":10023"
)

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	d := &DataServer{}
	pb.RegisterDataServer(s, d)
	s.Serve(lis)
}

type DataServer struct {
}

func (s *DataServer) GetUserInfo(ctx context.Context, ui *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	return nil, nil
}

func (s *DataServer) ChangeUserInfo(d pb.Data_ChangeUserInfoServer) error {
	return nil
}
