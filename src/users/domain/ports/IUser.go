package ports

import "gRPC-Example/src/users/domain/entities"

type IUser interface {
	Save(user *entities.User) error
	ListAll() ([]entities.User, error)
	Update(id int32, user *entities.User) error
	Delete(id int32) error
}