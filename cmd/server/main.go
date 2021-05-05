package main

import (
	"context"
	"log"
	"net"

	"github.com/iamthe1whoknocks/testTaskStrafovNet/pkg/parser"
	pb "github.com/iamthe1whoknocks/testTaskStrafovNet/proto/testTaskStrafovNet"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedAPIServer
}

func (g *GRPCServer) Get(ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	log.Println("Got request, INN = ", req.INN)
	resp, err = parser.GetCompanyInfo(req.INN)

	if err != nil {
		return nil, err
	}

	return resp, nil
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
