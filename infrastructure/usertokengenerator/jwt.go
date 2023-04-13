package usertokengenerator

import (
	"usermanagement/model"

	"github.com/golang-jwt/jwt/v5"
)

var (
	keyDefault = []byte("DEFAULT_JWT_KEY")
)

type jwtUTG struct {
	key []byte
}

func NewJWT() *jwtUTG {
	return &jwtUTG{
		key: keyDefault,
	}
}

func (j *jwtUTG) Generate(user model.User) (string, error) {
	// FIXME: Use custom claim to have the expired at
	// FIXME: Validate JWT in middleware
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	})
	return tkn.SignedString(j.key)
}
