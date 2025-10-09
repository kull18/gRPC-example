package infraestructure

import (
	"context"
	pb "gRPC-Example/proto"
	"gRPC-Example/src/application"
	"gRPC-Example/src/domain/entities"
)

type UpdateUserController struct {
	uc *application.UpdateUserUseCase
	pb.UnimplementedUpdateUserServer
}

func NewUpdateUserController(uc *application.UpdateUserUseCase) *UpdateUserController {
	return &UpdateUserController{
		uc: uc, 
	}
}

func (ctr *UpdateUserController) UpdateUser(context context.Context, r *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := &entities.User{
		Username: r.Username,
		Password: r.Password,
	}

	err := ctr.uc.Run(r.ID, user)

	if err != nil {
		return &pb.UpdateUserResponse{
			Message: "Cant update user",
			Success: false,
		}, nil
	}

	return &pb.UpdateUserResponse{
		Message: "Product updated",
		Success: true,
	}, nil
}