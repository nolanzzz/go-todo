package hash

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct {
	cost int
}

// Make encryption method
func (b *Bcrypt) Make(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, b.cost)
}

// Check validate given password
func (b *Bcrypt) Check(hashed, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashed, password)
}
