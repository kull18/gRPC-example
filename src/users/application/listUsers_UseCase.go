package application

import (
	"gRPC-Example/src/users/domain/entities"
	"gRPC-Example/src/users/domain/ports"
	"log"
)

type ListUsersUseCase struct {
	repository ports.IUser
}

func NewListUsersUseCase(repository ports.IUser) *ListUsersUseCase {
	return &ListUsersUseCase{
		repository: repository,
	}
}

func (uc *ListUsersUseCase) Run() ([]entities.User, error) {
	users, err := uc.repository.ListAll()

	if err != nil {
		log.Printf("error to list all users")
		return nil, err
	}

	return users, nil
}