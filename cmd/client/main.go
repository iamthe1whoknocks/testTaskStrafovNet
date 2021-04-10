package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/DimkaTheGreat/testTaskStrafovNet/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting client....")
	conn, err := grpc.Dial(":55555", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}

	c := pb.NewAPIClient(conn)

	req := &pb.Request{
		INN: "fgdgdfgdfgd",
	}

	resp, err := c.Get(context.Background(), req)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Printf("%s\n%s\n%s\n%s\n", resp.INN, resp.KPP, resp.Name, resp.Leader)

}
