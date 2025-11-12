package main

import (
	"context"
	"fmt"
	"net"

	grpcPb "github.com/hello/grpc/proto"
	"google.golang.org/grpc"
)

type GrpcServerType struct {
	grpcPb.UnimplementedEventsServer
}

func (s *GrpcServerType) GetEvents(ctx context.Context, req *grpcPb.InputGetEvents) (*grpcPb.Event, error) {

	return &grpcPb.Event{
		Name:        "test event",
		Description: "test event Description",
	}, nil

}

func main() {
	fmt.Println("[msg] >>>> staring grpc server...")
	listener, err := net.Listen("tcp", ":5811")
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
