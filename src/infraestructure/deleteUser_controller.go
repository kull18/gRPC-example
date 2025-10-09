package infraestructure

import (
	"context"
	"gRPC-Example/src/application"
	pb "gRPC-Example/proto"
)

type DeleteUserController struct {
	uc *application.DeleteUserUseCase
	pb.UnimplementedDeleteUserServer
}

func NewDeleteUserController(uc *application.DeleteUserUseCase) *DeleteUserController {
	return &DeleteUserController{
		uc: uc,  
	}
}

func (ctr *DeleteUserController) DeleteUser(context context.Context, r *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err :=  ctr.uc.Run(r.ID)

	if err != nil {
		return &pb.DeleteUserResponse{
			Success: false,
			Message:  "error to delete user" + err.Error(),
		}, err
	}

	return &pb.DeleteUserResponse{
		Success: true,
		Message: "User deleted",
	}, nil
}