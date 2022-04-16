package database

import (
	"github.com/gofrs/uuid"
	"github.com/h1ght1me/auth-micro/pkg/database/user"
)

type UserService interface {
	User(id uuid.UUID) (*user.User, error)
	Users() ([]*user.User, error)
	CreateUser(t *user.User) error
	DeleteUser(id uuid.UUID) error
}
