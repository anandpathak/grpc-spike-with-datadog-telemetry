package main
import (
  "fmt"
  "google.golang.org/grpc"
  "demogrpc/internal/grpc/service"
  "log"
  "net"
)

func main() {
  netListener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8001))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
    panic(fmt.Sprintf("failed to listen: %v", err))
  }
  grpcServer := grpc.NewServer()

  ping := service.PingService{}
  service.RegisterPingServiceServer(grpcServer, ping)

  if err := grpcServer.Serve(netListener); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }
  fmt.Println("gRPC In Action!")
}