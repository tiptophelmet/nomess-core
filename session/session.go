package session

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/tiptophelmet/nomess-core/v5/errs"
	"github.com/tiptophelmet/nomess-core/v5/logger"
)

type sessionManager struct {
	jwtExpTime   int64
	jwtExpWindow int64
	jwtSecret    string
}

var session *sessionManager

func Init(jwtExpTime, jwtExpWindow int64, jwtSecret string) {
	session = &sessionManager{jwtExpTime, jwtExpWindow, jwtSecret}
}

func Get() *sessionManager {
	if session == nil {
		logger.Fatal("session manager was not initialized")
	}

	return session
}

func (session *sessionManager) IssueSessionToken(userID string) (string, error) {
	// TODO: Add jti support
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Unix() + session.jwtExpTime,
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedJwt, err := token.SignedString([]byte(session.jwtSecret))
	if err != nil {
		logger.Fatal("failed to sign jwt: %v", err.Error())
		return "", errs.ErrJwtNotIssued
	}

	return signedJwt, nil
}

func (session *sessionManager) checkRotationCondition(jwtClaimedExpTime int64) (bool, error) {
	currentTime := time.Now().Unix()

	isExpired := currentTime > jwtClaimedExpTime
	isRotationWindowOk := currentTime < jwtClaimedExpTime+session.jwtExpWindow

	return isExpired && isRotationWindowOk, nil
}

func (session *sessionManager) TryRotateSessionToken(signedToken string) (string, error) {
	parsedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(session.jwtSecret), nil
	})

	if !parsedToken.Valid {
		return "", errs.ErrInvalidJwt
	}

	_, isTokenExpError := err.(*jwt.TokenExpiredError)

	if err != nil && !isTokenExpError {
		return "", errs.ErrJwtNotParsed
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errs.ErrInvalidJwtClaims
	}

	jwtClaimedExpTime := claims["exp"].(int64)

	isRotationOk, err := session.checkRotationCondition(jwtClaimedExpTime)
	if err != nil {
		return "", err
	}

	if isRotationOk {
		return session.IssueSessionToken(claims["sub"].(string))
	}

	return "", nil
}

func (session *sessionManager) ValidateSessionToken(token string) (bool, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(session.jwtSecret), nil
	})

	if !parsedToken.Valid || err != nil {
		return false, nil
	}

	return true, nil
}
