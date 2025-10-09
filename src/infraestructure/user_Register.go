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
	deleteUserController *DeleteUserController
}

func NewUserRegister(createUserController *CreateUserController, getAllUsersController *GetAllUsersController, updateUserController *UpdateUserController, deleteUserController *DeleteUserController,grpcServer *grpc.Server)  *UserRegister {
	return  &UserRegister{
		createUserController: createUserController,
		getAllUsersController: getAllUsersController,
		updateUserController: updateUserController,
		grpcServer: grpcServer,
		deleteUserController: deleteUserController,
	}
}


func (register *UserRegister) Register() {
	pb.RegisterUserCreateServiceServer(register.grpcServer, register.createUserController)
	pb.RegisterUserGetAllServiceServer(register.grpcServer, register.getAllUsersController)
	pb.RegisterUpdateUserServer(register.grpcServer, register.updateUserController)
	pb.RegisterDeleteUserServer(register.grpcServer, register.deleteUserController)
}