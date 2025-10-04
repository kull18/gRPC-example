package ports

import "gRPC-Example/src/domain/entities"

type IUser interface {
	Save(user *entities.User) error
}