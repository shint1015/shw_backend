package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	shwgrpc "shwgrpc/pkg/grpc"
)

type ShwServer struct {
	shwgrpc.UnimplementedHelloServiceServer
	shwgrpc.UnimplementedHouseworkServiceServer
	shwgrpc.UnimplementedFamilyServiceServer
	shwgrpc.UnimplementedUserServiceServer
}

func NewShwServer() *ShwServer {
	return &ShwServer{}
}

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	shwgrpc.RegisterHelloServiceServer(s, NewShwServer())

	reflection.Register(s)

	go func() {
		log.Printf("gRPC server is running on port %d", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down gRPC server...")
	s.GracefulStop()
}

func (s *ShwServer) Hello(ctx context.Context, req *shwgrpc.HelloRequest) (*shwgrpc.HelloResponse, error) {
	return &shwgrpc.HelloResponse{Message: "Hello " + req.Name}, nil
}
