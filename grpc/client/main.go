package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/hello/config"
	grpcPb "github.com/hello/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func printEvent(client grpcPb.EventsClient, input *grpcPb.InputGetEvents) {
	fmt.Println("[msg] >>>> EVENT InputGetEvents", input)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	event, err := client.GetEvents(ctx, input)
	if err != nil {
		fmt.Println("[err] >>>> fetch failed", err)
		return
	}
	fmt.Println("[msg] >>>>> success >>>>>", event)
}
func main() {
	var opts []grpc.DialOption
	config.LoadConfig()
	var grpcHostAddress = fmt.Sprintf(":%d", config.Config.GrpcPort)
	fmt.Println("host address ->", grpcHostAddress, config.Config)
	var serverAddr = flag.String("addr", grpcHostAddress, "[msg] >>>> The server address in the format of host:port")
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		fmt.Println("[err] >>>>", err)
	}
	defer conn.Close()
	client := grpcPb.NewEventsClient(conn)
	printEvent(client, &grpcPb.InputGetEvents{PageNo: 1, PageSize: 10})
}
