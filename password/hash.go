package password

import (
	"github.com/tiptophelmet/nomess-core/v5/errs"
	"github.com/tiptophelmet/nomess-core/v5/logger"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		logger.Fatal("failed to hash a password: %s", err.Error())
		return "", errs.ErrPasswordHash
	}
	return string(hash), nil
}
