package response

import (
	"fmt"
	"io"
	pb "rpcTest/rpcbuild/rpcbuild/friday"

	"golang.org/x/net/context"
)

type Server struct {
	routeNotes []*pb.UserInfoResponse
}

//簡單模式
func (this *Server) GetUserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	uid := in.GetUid()
	fmt.Println("The uid is ", uid)
	return &pb.UserInfoResponse{
		Name:  "Jim",
		Age:   18,
		Sex:   0,
		Count: 1000,
	}, nil
}

//雙向流模式
func (this *Server) ChangeUserInfo(stream pb.Data_ChangeUserInfoServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("read done")
			return nil
		}
		if err != nil {
			fmt.Println("ERR", err)
			return err
		}
		fmt.Println("userinfo ", in)
		for _, note := range this.routeNotes {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}
