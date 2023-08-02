package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	jwt.StandardClaims
	Username        string `json:"username"`
	ApplicationName string
}

const KEY = "abcdefghiklmnopqrstuABCDEFGHIKLMNOPQRSTUVWXYZ~"

func GenerateToken(username string) (string, error) {
	now := time.Now().UTC()
	end := now.Add(1 * time.Hour)

	claim := &JwtClaims{
		Username: username,
	}

	claim.IssuedAt = now.Unix()
	claim.ExpiresAt = end.Unix()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := t.SignedString([]byte(KEY))
	if err != nil {
		return "", fmt.Errorf("generate token error : %w", err)
	}
	return token, nil
}

func VerifyAccessToken(tokenString string) (string, error) {
	claim := &JwtClaims{}
	t, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(KEY), nil
	})

	if err != nil {
		return "", fmt.Errorf("verify token error : %w", err)
	}
	if !t.Valid {
		return "", fmt.Errorf("verify token error : Invalid token")
	}

	return claim.Username, nil
}
