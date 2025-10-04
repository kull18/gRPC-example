package application

import (
	"gRPC-Example/src/domain/entities"
	"gRPC-Example/src/domain/ports"
)

type SaveUserUseCase struct {
	IUser ports.IUser
}


func NewSaveUserUseCase(iUser ports.IUser) *SaveUserUseCase {
	return &SaveUserUseCase{
		IUser: iUser,
	}
}

func (uc *SaveUserUseCase)  run(user *entities.User) error {
	err := uc.IUser.Save(user)

	if err != nil {
		return err
	}

	return nil
}