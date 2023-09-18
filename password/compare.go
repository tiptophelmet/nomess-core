package password

import (
	"github.com/tiptophelmet/nomess-core/v2/logger"

	"golang.org/x/crypto/bcrypt"
)

func CompareToHash(password string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false
	} else if err != nil {
		logger.Fatal("failed to compare password to hash: %s", err.Error())
		return false
	}

	return true
}
