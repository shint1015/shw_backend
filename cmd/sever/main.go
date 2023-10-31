package main

import (
	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"context"
	"fmt"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"shwgrpc/http/controller"
	shwgrpc "shwgrpc/pkg/grpc"
	"shwgrpc/pkg/grpc/grpcconnect"
)

var houseworkController = controller.NewHouseworkController()
var familyController = controller.NewFamilyController()
var userController = controller.NewUserController()

type ShwServer struct {
	//shwgrpc.UnimplementedHelloServiceServer
	//shwgrpc.UnimplementedHouseworkServiceServer
	//shwgrpc.UnimplementedFamilyServiceServer
	//shwgrpc.UnimplementedUserServiceServer
	grpcconnect.UnimplementedHelloServiceHandler
	grpcconnect.UnimplementedHouseworkServiceHandler
	grpcconnect.UnimplementedFamilyServiceHandler
	grpcconnect.UnimplementedUserServiceHandler
}

func NewShwServer() *ShwServer {
	return &ShwServer{}
}

func newServeMuxWithReflection() *http.ServeMux {
	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		"shw.HelloService",
		"shw.HouseworkService",
		"shw.FamilyService",
		"shw.UserService",
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	return mux
}

func newInterceptors() connect.Option {
	interceptors := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			req.Header().Set("hoge", "fuga")
			return next(ctx, req)
		})
	}
	return connect.WithInterceptors(connect.UnaryInterceptorFunc(interceptors))
}

func main() {
	//fmt.Println("Hello World")
	//domain := "localhost"
	port := 8080
	shw := NewShwServer()
	//mux := newServeMuxWithReflection()
	interceptor := newInterceptors()
	mux := newServeMuxWithReflection()
	mux.Handle(grpcconnect.NewFamilyServiceHandler(shw, interceptor))
	mux.Handle(grpcconnect.NewHouseworkServiceHandler(shw, interceptor))
	mux.Handle(grpcconnect.NewUserServiceHandler(shw, interceptor))
	mux.Handle(grpcconnect.NewHelloServiceHandler(shw, interceptor))

	//mux.Handle(NewHouse)
	log.Printf("gRPC server is running on port %d", port)
	corsHandler := makeCorsHandler()
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		corsHandler.Handler(h2c.NewHandler(mux, &http2.Server{})),
	)
	if err != nil {
		fmt.Println(err)
	}

	log.Println("Shutting down gRPC server...")

	//listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//s := grpc.NewServer()
	//
	//shwgrpc.RegisterHelloServiceServer(s, NewShwServer())
	//shwgrpc.RegisterFamilyServiceServer(s, NewShwServer())
	//shwgrpc.RegisterHouseworkServiceServer(s, NewShwServer())
	//shwgrpc.RegisterUserServiceServer(s, NewShwServer())
	//
	//reflection.Register(s)
	//
	//go func() {
	//	log.Printf("gRPC server is running on port %d", port)
	//	s.Serve(listener)
	//}()
	//
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, os.Interrupt)
	//<-quit
	//
	//log.Println("Shutting down gRPC server...")
	//s.GracefulStop()
}

func makeCorsHandler() *cors.Cors {
	// TODO: Originを環境変数から取得する
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://172.19.0.2:3000", "http://localhost", "http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})
	return corsHandler
}

func (s *ShwServer) Hello(ctx context.Context, req *connect.Request[shwgrpc.HelloRequest]) (*connect.Response[shwgrpc.HelloResponse], error) {
	res := connect.NewResponse(&shwgrpc.HelloResponse{Message: "Hello " + req.Msg.Name})
	return res, nil
}
