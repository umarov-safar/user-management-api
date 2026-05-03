package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/umarov-safar/user-management-api/internal/models"
)

var ErrInvalidToken = errors.New("invalid token")

type JWToken struct {
	secret     string
	expiration int64
}

func (j *JWToken) SetSecretKey(secret string) *JWToken {
	j.secret = secret
	return j
}

func (j *JWToken) SetExpiration(expiration int64) *JWToken {
	j.expiration = expiration
	return j
}

func (j *JWToken) GetSecretKey() string {
	return j.secret
}

func (j *JWToken) GetExpiration() int64 {
	return j.expiration
}

func (j *JWToken) Generate(userId, email, role, secret string, expiration int64) (string, error) {
	claims := models.CustomClaims{
		UserID: userId,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expiration))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (j *JWToken) ParseJWT(tokenStr, secret string) (*models.CustomClaims, error) {
	claims := &models.CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
