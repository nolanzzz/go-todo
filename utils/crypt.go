package utils

import "todo/core/hash"

func CreatePass(pass string) (password string) {
	hashed, _ := hash.NewHash().Make([]byte(pass))
	password = string(hashed)
	return password
}
