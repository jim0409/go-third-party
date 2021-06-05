package mock_practice_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	pb "go-third-party/gRPC/practice_server/proto"

	// hwmock "go-third-party/gRPC/hello_server/mock_unittest"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	pbmock "go-third-party/gRPC/practice_server/mock_unittest"
)

// rpcMsg implements the gomock.Matcher interface
type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

func TestCheckHealth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGreeterClient := pbmock.NewMockHealthClient(ctrl)
	// req := &helloworld.HelloRequest{Name: "unit_test"}
	req := &pb.CheckHealthRequest{
		Personinfo: &pb.PersonInfo{
			Name: "jim",
		},
	}
	mockGreeterClient.EXPECT().CheckHealth(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&pb.CheckHealthReply{
		Reportinfo: &pb.ReportInfo{
			Message: "Mocked Interface",
		},
	}, nil)
	testCheckHealth(t, mockGreeterClient)
}

// func testSayHello(t *testing.T, client helloworld.GreeterClient) {
func testCheckHealth(t *testing.T, client pb.HealthClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.CheckHealth(ctx, &pb.CheckHealthRequest{
		Personinfo: &pb.PersonInfo{
			Name: "jim",
		},
	})
	if err != nil || r.Reportinfo.Message != "Mocked Interface" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.Reportinfo.Message)
}
