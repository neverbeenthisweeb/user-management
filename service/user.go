package service

import (
	"context"
	"usermanagement/infrastructure"
	"usermanagement/model"
	"usermanagement/repository"

	"github.com/rs/zerolog/log"
)

type User interface {
	Login(context.Context, model.User) (model.Token, error)
	List(context.Context, model.UserListFilter) ([]model.User, error)
	// AddUser(context.Context, model.User) (model.User, error)
	// RemoveUser(context.Context, model.User) error
}

type userImpl struct {
	repoUser repository.User

	infraHasher      infrastructure.Hasher
	infraUserTokenGn infrastructure.UserTokenGenerator
}

func NewUserImpl(infra *infrastructure.Infrastructure, repo *repository.Repository) *userImpl {
	return &userImpl{
		repoUser:         repo.User,
		infraHasher:      infra.Hasher,
		infraUserTokenGn: infra.UserTokenGenerator,
	}
}

func (s *userImpl) Login(ctx context.Context, m model.User) (model.Token, error) {
	logger := log.With().
		Str("method", "userImpl.Login").
		Str("requestid", ctx.Value("requestid").(string)).
		Logger()

	// Validation
	if m.Username == "" {
		logger.Error().Err(ErrMissingUsername).Msg("Missing username")
		return model.Token{}, ErrMissingUsername
	}

	if m.Password == "" {
		logger.Error().Err(ErrMissingPassword).Msg("Missing password")
		return model.Token{}, ErrMissingPassword
	}

	// Get existing password from store
	existingPw, err := s.repoUser.GetUserByUsername(ctx, m.Username)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get password")
		return model.Token{}, err
	}

	// Compare password

	// debugPw, err := s.infraHasher.GenerateFromPassword([]byte(m.Password), -1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(debugPw)) // FIXME: Debugging to create hash

	err = s.infraHasher.CompareHashAndPassword([]byte(existingPw), []byte(m.Password))
	if err != nil {
		logger.Error().Err(err).Msg("Password does not match")
		return model.Token{}, ErrPasswordWrong
	}

	// Generate token
	token, err := s.infraUserTokenGn.Generate(model.User{Username: m.Username})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to generate token")
		return model.Token{}, err
	}

	return model.Token{AccessToken: token}, nil
}

func (u *userImpl) List(ctx context.Context, filter model.UserListFilter) ([]model.User, error) {
	logger := log.With().
		Str("method", "userImpl.List").
		Str("requestid", ctx.Value("requestid").(string)).
		Logger()

	// Get list of users
	users, err := u.repoUser.FetchUser(ctx, filter)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch users")
		return []model.User{}, err
	}

	return users, nil
}
