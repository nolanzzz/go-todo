package hash

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestNewHash(t *testing.T) {
	hash := Bcrypt{
		cost: bcrypt.DefaultCost,
	}
	password := "123456"
	bytes, err := hash.Make([]byte(password))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(bytes))

	err = hash.Check(bytes, []byte(password))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Correct!")
}
