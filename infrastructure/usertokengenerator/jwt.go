package usertokengenerator

import (
	"time"
	"usermanagement/model"

	"github.com/golang-jwt/jwt/v5"
)

var (
	KeyDefault = []byte("DEFAULT_JWT_KEY")
	expiryDur  = 5 * time.Minute
)

type claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type jwtUTG struct {
	key []byte
}

func NewJWT() *jwtUTG {
	return &jwtUTG{
		key: KeyDefault,
	}
}

func (j *jwtUTG) Generate(user model.User) (string, error) {
	cl := claims{
		user.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiryDur)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, cl)
	return tkn.SignedString(j.key)
}
