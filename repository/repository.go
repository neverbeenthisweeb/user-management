package repository

import "context"

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
	// Get password of an active user
	GetPassword(ctx context.Context, username string) (string, error)
}
