package password

import (
	"github.com/tiptophelmet/nomess-core/errs"
	"github.com/tiptophelmet/nomess-core/logger"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		logger.Fatal("could not hash a password: %s", err.Error())
		return "", errs.ErrPasswordHash
	}
	return string(hash), nil
}
