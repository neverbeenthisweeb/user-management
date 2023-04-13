package inmemory

import (
	"context"
	"usermanagement/model"
	"usermanagement/service"
)

var (
	// FIXME: Move to JSON?
	// FIXME: Init in constructor?
	userDB = []model.User{
		{
			ID:        "ID_1",
			Username:  "USERNAME_1",
			Password:  "PASSWORD_1", // FIXME: Replace with bcrypt
			IsDeleted: false,
		},
		{
			ID:        "ID_2",
			Username:  "USERNAME_2",
			Password:  "PASSWORD_2", // FIXME: Replace with bcrypt
			IsDeleted: true,
		},
	}
)

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) GetPassword(ctx context.Context, username string) (string, error) {
	for _, v := range userDB {
		if v.Username == username && !v.IsDeleted {
			return v.Password, nil
		}
	}

	return "", service.ErrPasswordNotFound
}
