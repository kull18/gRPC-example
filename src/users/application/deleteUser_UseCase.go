package application

import "gRPC-Example/src/users/domain/ports"

type DeleteUserUseCase struct {
	repository ports.IUser
}

func NewDeleteUserUseCase(repository ports.IUser) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		repository: repository,
	}
}

func (uc *DeleteUserUseCase) Run(id int32) error {
	err := uc.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}