package dao

import "github.com/golang-jwt/jwt/v5"

type MainClaims struct {
	Uid  string `json:"uid"`
	Role string `json:"role"`
	Type string `json:"type"`

	jwt.RegisteredClaims
}
