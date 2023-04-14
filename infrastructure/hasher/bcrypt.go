package hasher

import "golang.org/x/crypto/bcrypt"

type hasherImpl struct{}

func NewDefaultHasher() *hasherImpl {
	return &hasherImpl{}
}

func (h *hasherImpl) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func (h *hasherImpl) Cost(hashedPassword []byte) (int, error) {
	return bcrypt.Cost(hashedPassword)
}

func (h *hasherImpl) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}
