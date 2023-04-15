package entity

import "github.com/golang-jwt/jwt"

type SSJWTClaim struct {
	UserID int64 `json:"user_id"`
	*jwt.StandardClaims
}
