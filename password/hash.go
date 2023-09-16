package password

import (
	"github.com/tiptophelmet/nomess-core/v2/errs"
	"github.com/tiptophelmet/nomess-core/v2/logger"

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
