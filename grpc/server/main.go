package main

import (
	"context"
	"fmt"
	"net"

	"github.com/hello/config"
	"github.com/hello/databases"
	grpcPb "github.com/hello/grpc/proto"
	"google.golang.org/grpc"
)

type GrpcServerType struct {
	grpcPb.UnimplementedEventsServer
}

func initServer() {
	config.LoadConfig()
	databases.InitDb()
}

func (s *GrpcServerType) GetEvents(ctx context.Context, req *grpcPb.InputGetEvents) (*grpcPb.Event, error) {

	return &grpcPb.Event{
		Name:        "test event",
		Description: "test event Description",
	}, nil

}

func main() {
	fmt.Println("[msg] >>>> staring grpc server...")
	initServer()
	grpcHostAddress := fmt.Sprintf(":%d", config.Config.GrpcPort)
	listener, err := net.Listen("tcp", grpcHostAddress)
	if err != nil {
		fmt.Println("[err] >>>>> Error with starting GRPC server", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	grpcPb.RegisterEventsServer(grpcServer, &GrpcServerType{})

	// serve
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Println("[err] >>>>> Error in serving GRPC request", err)
	}
	fmt.Println("[msg] >>>>> ")
}
