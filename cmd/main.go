package main

import (
	"gRPC-Example/src/application"
	"gRPC-Example/src/infraestructure"
	"log"
	"net"

	pb "gRPC-Example/proto"

	"google.golang.org/grpc"
)

func main() {
	adapter := infraestructure.NewMysqlAdapter()
	uc := application.NewSaveUserUseCase(adapter)
	ctr := infraestructure.NewCreateUserController(uc)

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {

	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, ctr)


	log.Print("Listen server grpc")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve", err)
	}
}