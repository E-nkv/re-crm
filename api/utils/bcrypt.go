package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Bcryptify(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("😡 err bcrpytifing pass", err)
	}
	return string(hash)
}
