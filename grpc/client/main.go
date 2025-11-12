package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	grpcPb "github.com/hello/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverAddr = flag.String("addr", "localhost:5811", "[msg] >>>> The server address in the format of host:port")

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
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		fmt.Println("[err] >>>>", err)
	}
	defer conn.Close()
	client := grpcPb.NewEventsClient(conn)
	printEvent(client, &grpcPb.InputGetEvents{PageNo: 1, PageSize: 10})
}
