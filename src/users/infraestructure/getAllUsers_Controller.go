package infraestructure

import (
	"context"
	pb "gRPC-Example/proto"
	"gRPC-Example/src/users/application"
	"gRPC-Example/src/users/domain/entities"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

type GetAllUsersController struct {
	uc *application.ListUsersUseCase
	pb.UnimplementedUserGetAllServiceServer
}

func NewGetAllUsersController(uc *application.ListUsersUseCase) *GetAllUsersController {
	return &GetAllUsersController{
		uc: uc, 
	}
}

func convertToProtoUser(users []entities.User) []*pb.User {
	protoUsers := make([]*pb.User, 0, len(users))

	for _, u := range users {
		protoUsers = append(protoUsers, &pb.User{
			ID: int32(u.ID),
			Username: u.Username,
			Password: u.Password,
		})
	}

	return protoUsers
}

func (ctr *GetAllUsersController) GetUsers(context context.Context, empty *emptypb.Empty) (*pb.ListUserResponse, error){
	users, err := ctr.uc.Run()

	if err != nil {
		log.Printf("error to list users")
		return nil, err
	}

	protoUsers := convertToProtoUser(users)

	return &pb.ListUserResponse{
		Users: protoUsers,
	}, nil
}