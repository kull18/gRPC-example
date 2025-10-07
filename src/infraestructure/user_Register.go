package infraestructure

import (
	pb "gRPC-Example/proto"

	"google.golang.org/grpc"
)

type UserRegister struct {
	grpcServer *grpc.Server
	createUserController *CreateUserController
	getAllUsersController *GetAllUsersController
	updateUserController *UpdateUserController
}

func NewUserRegister(createUserController *CreateUserController, getAllUsersController *GetAllUsersController, updateUserController *UpdateUserController,grpcServer *grpc.Server)  *UserRegister {
	return  &UserRegister{
		createUserController: createUserController,
		getAllUsersController: getAllUsersController,
		updateUserController: updateUserController,
		grpcServer: grpcServer,
	}
}


func (register *UserRegister) Register() {
	pb.RegisterUserCreateServiceServer(register.grpcServer, register.createUserController)
	pb.RegisterUserGetAllServiceServer(register.grpcServer, register.getAllUsersController)
	pb.RegisterUpdateUserServer(register.grpcServer, register.updateUserController)
}