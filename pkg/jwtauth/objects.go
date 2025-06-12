package jwtauth

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	Role          string `json:"role,omitempty"`
	IsPartial     bool   `json:"is_partial,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
	PhoneNumber   string `json:"phone_number,omitempty"`
	Email         string `json:"email,omitempty"`
	AccountUUID   string `json:"account_uuid,omitempty"`
}

type JWTTokenClaims struct {
	CustomClaims
	jwt.RegisteredClaims
}
