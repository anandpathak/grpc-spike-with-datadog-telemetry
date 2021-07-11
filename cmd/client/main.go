package main

import (
	"context"
	"demogrpc/internal/grpc/domain"
	"demogrpc/internal/grpc/service"
	"demogrpc/internal/telemetry"
	"sync"
	"time"

	"go.opencensus.io/plugin/ocgrpc"

	"fmt"

	"google.golang.org/grpc"
)

func main() {
	telemetry.Telemetry()
	serverAddress := "localhost:8001"
	var wg sync.WaitGroup
	wg.Add(1)
	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithStatsHandler(new(ocgrpc.ClientHandler)))

	if e != nil {
		panic(e)
	}

	defer conn.Close()

	client := service.NewPingServiceClient(conn)
	for _ = range [10]int{} {
		repositoryModel := domain.PingMessage{
			Test: "hey",
		}
		if responseMessage, e := client.Ping(context.Background(), &repositoryModel); e != nil {
			panic(fmt.Sprintf("Was not able to insert Record %v", e))
		} else {
			fmt.Println(responseMessage)
		}
	}
	time.Sleep(60*time.Second)
	wg.Done()
}
