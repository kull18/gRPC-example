package ports

import "gRPC-Example/src/domain/entities"

type IUser interface {
	Save(user *entities.User) error
	ListAll() ([]entities.User, error)
}