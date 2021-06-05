package main

import (
	"context"
	"log"
	"net"

	pb "go-third-party/gRPC/practice_server/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) CheckHealth(ctx context.Context, in *pb.CheckHealthRequest) (*pb.CheckHealthReply, error) {
	log.Printf("Received: %v's data", in.Personinfo.Name)
	// fmt.Printf("The request data includes %v\n", in)

	healthReport := returnHealthReport(in)
	// fmt.Printf("The healthReport data includes %v\n", healthReport)

	return &pb.CheckHealthReply{
		Reportinfo: healthReport,
	}, nil
}

func returnHealthReport(in *pb.CheckHealthRequest) *pb.ReportInfo {
	// fmt.Printf("The request data includes %v\n", in)
	var healthReport = &pb.ReportInfo{}
	msg := ""
	isHealth := false
	bmi := in.Personinfo.Weight * 10000 / (in.Personinfo.Hight * in.Personinfo.Hight)
	if bmi > 25 || bmi < 15 {
		msg = "You are not health"
	} else {
		msg = "You are very health"
		isHealth = true
	}
	healthReport = &pb.ReportInfo{
		Name:     in.Personinfo.Name,
		Message:  msg,
		BMI:      bmi,
		IsHealth: isHealth,
	}
	return healthReport
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHealthServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
