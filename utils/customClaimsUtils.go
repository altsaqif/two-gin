package utils

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	jwt.RegisteredClaims
	Role     string `json:"role"`
	AuthorId string `json:"authorId"`
}
