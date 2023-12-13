package repository

import (
	"secret-santa/pkg"

	"github.com/jmoiron/sqlx"
)

type Authentication interface {
	IsUserExists(user pkg.User) (bool, error)
	AddUser(user pkg.User) (int, error)
}

type Room interface {
}

type Repository struct {
	Authentication
	Room
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authentication: sqlite.,
	}
}
