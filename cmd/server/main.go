package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"shwgrpc/config"
	"shwgrpc/http/controller"
	"shwgrpc/internal/auth"
	"shwgrpc/model"
	shwgrpc "shwgrpc/pkg/grpc"
	"shwgrpc/pkg/grpc/grpcconnect"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var houseworkController = controller.NewHouseworkController()
var familyController = controller.NewFamilyController()
var userController = controller.NewUserController()
var pointController = controller.NewPointController()

type ShwServer struct {
	grpcconnect.UnimplementedHelloServiceHandler
	grpcconnect.UnimplementedHouseworkServiceHandler
	grpcconnect.UnimplementedFamilyServiceHandler
	grpcconnect.UnimplementedUserServiceHandler
	grpcconnect.UnimplementedPointServiceHandler
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
		"shw.PointService",
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	return mux
}

func newInterceptors() connect.Option {
	validator, err := auth.NewAuth0ValidatorFromEnv()
	if err != nil {
		log.Fatalf("auth0 config error: %v", err)
	}
	authInterceptor := auth.NewInterceptor(validator, []string{
		"/shw.HelloService/Hello",
	})
	return connect.WithInterceptors(authInterceptor.Unary())
}

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Error loading env: %v", err)
	}
	model.Init()
	//fmt.Println("Hello World")
	//domain := "localhost"
	port := 8080
	shw := NewShwServer()
	interceptor := newInterceptors()
	mux := newServeMuxWithReflection()
	mux.Handle(grpcconnect.NewFamilyServiceHandler(shw, interceptor))
	mux.Handle(grpcconnect.NewHouseworkServiceHandler(shw, interceptor))
	mux.Handle(grpcconnect.NewUserServiceHandler(shw, interceptor))
	mux.Handle(grpcconnect.NewHelloServiceHandler(shw, interceptor))
	mux.Handle(grpcconnect.NewPointServiceHandler(shw, interceptor))

	//mux.Handle(NewHouse)
	log.Printf("gRPC server is running on port %d", port)
	corsHandler := makeCorsHandler()
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		corsHandler.Handler(h2c.NewHandler(mux, &http2.Server{})),
	)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Shutting down gRPC server...")
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
