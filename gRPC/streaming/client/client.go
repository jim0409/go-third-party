package main

import (
    "net"
    "google.golang.org/grpc"
    pb "rpcTest/rpcbuild/rpcbuild/friday"
    "rpcTest/rpcbuild/response"
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
    pb.RegisterDataServer(s, &response.Server{})
    s.Serve(lis)

}