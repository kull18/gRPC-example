package infraestructure

import (
	"context"
	pb "gRPC-Example/proto"
	"gRPC-Example/src/application"
	"gRPC-Example/src/domain/entities"
)

type CreateUserController struct {
	uc *application.SaveUserUseCase
	pb.UnimplementedUserCreateServiceServer
}

func NewCreateUserController(uc *application.SaveUserUseCase) *CreateUserController {
	return &CreateUserController{
		uc: uc,
	}
}

func (ctr *CreateUserController) SaveUser(ctx context.Context, req *pb.SaveUserRequest) (*pb.SaveUserResponse, error) {
	user := &entities.User{
		Username: req.Username,
		Password: req.Password,
	}

	err := ctr.uc.IUser.Save(user)
	if err != nil {
		return &pb.SaveUserResponse{
			Response: false,
			Message:  "Error saving user: " + err.Error(),
		}, nil
	}

	return &pb.SaveUserResponse{
		Response: true,
		Message:  "User created successfully",
	}, nil
}
