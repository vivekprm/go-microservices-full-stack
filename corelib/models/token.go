package models

import "github.com/golang-jwt/jwt"

type Token struct {
	UserID         string
	Name           string
	Email          string
	StandardClaims *jwt.StandardClaims
}

func (t *Token) Valid() error {
	return nil
}
