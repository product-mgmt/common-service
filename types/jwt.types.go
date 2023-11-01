package types

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID int `json:"userID"`
	jwt.RegisteredClaims
}
