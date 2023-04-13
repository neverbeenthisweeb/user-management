package infrastructure

import "usermanagement/model"

type Infrastructure struct {
	Hasher             Hasher // Password hasher
	UserTokenGenerator UserTokenGenerator
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{}
}

func (inf *Infrastructure) SetHasher(h Hasher) {
	inf.Hasher = h
}

func (inf *Infrastructure) SetUserTokenGenerator(utg UserTokenGenerator) {
	inf.UserTokenGenerator = utg
}

// FIXME: Create bcrypt hasher implementation
type Hasher interface {
	Hash(text []byte) ([]byte, error)
}

type UserTokenGenerator interface {
	Generate(user model.User) (string, error)
}
