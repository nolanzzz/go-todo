package jwt_helper

import (
	"log"
	"testing"
)

func TestDecode(t *testing.T) {

	decode, err := Decode("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwid2lkIjoiMCIsImV4cCI6MTY0Mzg2NTEwMn0.MqdZjQnootia6gORP7Q7Owtk_v3bCR0hTk40GRI1YKg")
	log.Println(decode.Wid)
	log.Println(decode.Username)
	if err != nil {
		log.Println(err.Error())
	}
}
