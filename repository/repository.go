package repository

import (
	"context"
	"usermanagement/model"
)

type Repository struct {
	User User
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) SetUser(u User) {
	r.User = u
}

type User interface {
	// Get an active user
	GetUserByUsername(ctx context.Context, username string) (string, error)
	FetchUser(ctx context.Context, filter model.UserListFilter) ([]model.User, error)
}
