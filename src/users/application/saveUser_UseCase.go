package application

import (
	"gRPC-Example/src/users/domain/entities"
	"gRPC-Example/src/users/domain/ports"
)

type SaveUserUseCase struct {
	IUser ports.IUser
}


func NewSaveUserUseCase(iUser ports.IUser) *SaveUserUseCase {
	return &SaveUserUseCase{
		IUser: iUser,
	}
}

func (uc *SaveUserUseCase)  Run(user *entities.User) error {
	err := uc.IUser.Save(user)

	if err != nil {
		return err
	}

	return nil
}