package jsondb

import (
	"context"
	"fmt"
	"usermanagement/model"
	"usermanagement/service"

	"github.com/sonyarouje/simdb"
)

type userRepository struct {
	db *simdb.Driver
}

func NewUserRepository() *userRepository {
	db, err := simdb.New("db")
	if err != nil {
		panic("Failed to init db: " + err.Error())
	}

	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (string, error) {
	var u model.User

	err := r.db.Open(User{}).
		Where("username", "=", username).
		Where("is_deleted", "=", false).
		First().
		AsEntity(&u)
	if err != nil {
		fmt.Println(err)
		if err == simdb.ErrRecordNotFound {
			return "", service.ErrPasswordNotFound
		}
		return "", err
	}

	return u.Password, nil
}

func (r *userRepository) FetchUser(ctx context.Context, filter model.UserListFilter) ([]model.User, error) {
	var uu []model.User

	q := r.db.Open(User{})

	if !filter.ShowDeleted {
		q = q.Where("is_deleted", "=", false)
	}

	err := q.Get().
		AsEntity(&uu)
	if err != nil {
		return nil, err
	}

	return uu, nil
}
