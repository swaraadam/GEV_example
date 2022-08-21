package main

import (
	"fmt"
	database "gev_example/database/connection"
	"gev_example/helpers"
	"gev_example/helpers/auth"
	"gev_example/models/protobuff"
	"gev_example/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// initialize database
	database.PG.InitializeDB()

	// create tcp connection on port 5555
	lis, err := net.Listen("tcp", helpers.Env.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to listen on address %v: %v", helpers.Env.ServerAddress, err.Error())
	}

	// create grpc server
	interceptor := auth.NewAuthInterceptor()
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	u := services.UserService{}
	protobuff.RegisterUserServiceServer(grpcServer, &u)
	reflection.Register(grpcServer)

	fmt.Println("service running on ", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over address %v: %v", helpers.Env.ServerAddress, err.Error())
	}
}
