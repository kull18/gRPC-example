package main

import (
	"gRPC-Example/src/application"
	"gRPC-Example/src/infraestructure"
	"log"
	"net"
	"google.golang.org/grpc"
)

func InitDependencies() {
	//here is where i declare 

	adapter := infraestructure.NewMysqlAdapter()
	uc := application.NewSaveUserUseCase(adapter)
	ctr := infraestructure.NewCreateUserController(uc)

	ucList := application.NewListUsersUseCase(adapter)
	ctrList := infraestructure.NewGetAllUsersController(ucList)

	ucUpdateUser := application.NewUpdateUserUseCase(adapter)
	ctrUpdateUser := infraestructure.NewUpdateUserController(ucUpdateUser)

	ucDeleteUser := application.NewDeleteUserUseCase(adapter)
	ctrDeleteUser := infraestructure.NewDeleteUserController(ucDeleteUser)

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Printf("error to init server")
	}

	grpcServer := grpc.NewServer()

	userRegister := infraestructure.NewUserRegister(ctr, ctrList,  ctrUpdateUser, ctrDeleteUser,grpcServer)
	userRegister.Register()

	log.Print("Listen server grpc")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve", err)
	}
}