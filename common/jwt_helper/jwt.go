package jwt_helper

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	key    []byte
	expire int64
)

type Claims struct {
	Username string `json:"username"`
	Wid      string `json:"wid"`
	jwt.StandardClaims
}

func init() {
	key = []byte("sockstack")
	expire = 7200
}

func Encode(c Claims) (string, error) {
	if c.ExpiresAt == 0 {
		c.ExpiresAt = time.Now().Unix() + expire
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Sign and get the complete encoded token as a string using the secret
	//if keys != nil {
	//	key = keys
	//}
	return token.SignedString(key)
}

func Decode(s string) (*Claims, error) {
	var err error
	// sample token is expired. override time so it parses as valid
	//if keys != nil {
	//	key = keys
	//}
	if s == "" {
		return &Claims{}, errors.New("empty token")
	}
	token, err := jwt.ParseWithClaims(s, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return key, nil
	})
	if err != nil {
		return &Claims{}, err
	}
	if !token.Valid {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				err = errors.New("That's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				err = errors.New("Timing is everything")
			} else {
				err = errors.New("Couldn't handle this token:")
			}
		} else {
			err = errors.New("Couldn't handle this token:")
		}
		return &Claims{}, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return &Claims{}, errors.New("Couldn't handle this token:")
	}
	return claims, nil
}
