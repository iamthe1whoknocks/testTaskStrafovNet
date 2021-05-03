package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	pb "github.com/DimkaTheGreat/testTaskStrafovNet/proto/testTaskStrafovNet"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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
		INN: "1656002652",
	}

	resp, err := c.Get(context.Background(), req)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Printf("%s\n%s\n%s\n%s\n", resp.INN, resp.KPP, resp.Name, resp.Leader)

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = pb.RegisterAPIHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":55556",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:55556")
	log.Fatalln(gwServer.ListenAndServe())
}

//curl -X POST -k http://localhost:55556/v1/post -d '{"INN": "1656002652"}'

//2913003750
