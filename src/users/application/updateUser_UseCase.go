package application

import (
	"gRPC-Example/src/users/domain/entities"
	"gRPC-Example/src/users/domain/ports"
)

type UpdateUserUseCase struct {
	repository ports.IUser
}

func NewUpdateUserUseCase(repository ports.IUser) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		repository: repository,
	}
}

func (uc *UpdateUserUseCase) Run(id int32, user *entities.User) error {
	err := uc.repository.Update(id, user)

	if err != nil {
		return err
	}
	
	return nil
}