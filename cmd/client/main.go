package main

import (
	"context"
	"log"

	pb "github.com/DimkaTheGreat/testTaskStrafovNet/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":55555", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}

	c := pb.NewAPIClient(conn)

	req := &pb.Request{
		INN: "34342",
	}

	resp, err := c.Get(context.Background(), req)

	log.Println(resp)

}
