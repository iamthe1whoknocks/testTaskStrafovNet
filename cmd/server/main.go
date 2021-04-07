package main

import (
	"context"
	"log"
	"net"

	pb "github.com/DimkaTheGreat/testTaskStrafovNet/pkg/api"
	"google.golang.org/grpc"
)

type GRPCServer struct{}

func (g *GRPCServer) Get(ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	return &pb.Response{
		INN:    req.GetINN(),
		KPP:    "kpp",
		Name:   "name",
		Leader: "leader",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":55555")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("start listening at port 55555")
	server := grpc.NewServer()
	grpcServer := &GRPCServer{}
	pb.RegisterAPIServer(server, grpcServer)

	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}