package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "go-third-party/gRPC/practice_server/proto"
	"google.golang.org/grpc"
)

var (
	name   = flag.String("name", "Testman", "for health check name")
	sexual = flag.String("sex", "NoSexual", "for health check sexual")
	hight  = flag.Int("hight", 180, "for health check hight")
	weight = flag.Int("weight", 75, "for health check weight")
)

const (
	address = "localhost:50051"
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHealthClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// r, err := c.SayHello(ctx, &pb.PracticeRequest{Name: name})
	requestCtx := pb.PersonInfo{
		Name:   *name,
		Sexual: *sexual,
		Hight:  int32(*hight),
		Weight: int32(*weight),
	}
	r, err := c.CheckHealth(ctx, &pb.CheckHealthRequest{
		Personinfo: &requestCtx,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(returnTemplate(r))
	// log.Printf("Greeting: %s", r.Reportinfo.Name)
}

func returnTemplate(r *pb.CheckHealthReply) string {
	return fmt.Sprintf("Welcome to check health %s, here is your health report \n===\nTo be brief: %s \nthe BMI is %d\nthe health cond is %v\n", r.Reportinfo.Name, r.Reportinfo.Message, r.Reportinfo.BMI, r.Reportinfo.IsHealth)
}
