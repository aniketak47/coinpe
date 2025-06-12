package controllers

import "coinpe/pkg/jwtauth"

type (
	TokenType string
	GoTo      string
)

const (
	GoToVerifyAccount  GoTo = "VERIFY_ACCOUNT"
	GotToCreateAccount GoTo = "CREATE_ACCOUNT"
	GoToContinue       GoTo = "CONTINUE"
)

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

type CreateTokenRequest struct {
	TokenType TokenType
	*jwtauth.CustomClaims
	ExpiryTimeInSeconds *int32
}

type CreateTokenResponse struct {
	Token           string
	ExpiryInSeconds int
}

type AuthenticateRequest struct {
	Username string `json:"username" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type AuthenticateResponse struct {
	AccessToken                 string `json:"access_token,omitempty"`
	RefreshToken                string `json:"refresh_token,omitempty"`
	AccessTokenExpiryInSeconds  int32  `json:"access_token_expiry_in_seconds,omitempty"`
	RefreshTokenExpiryInSeconds int32  `json:"refresh_token_expiry_in_seconds,omitempty"`
	IsNewUser                   bool   `json:"is_new_user,omitempty"`
	GoTo                        GoTo   `json:"goto,omitempty"`
	VerificationChannel         string `json:"verification_channel,omitempty"`
	Handle                      string `json:"handle,omitempty"`
}

type VerifyAuthRequest struct {
	Otp                        string `json:"otp,omitempty"`
	AccessToken                string `json:"access_token,omitempty"`
	AccessTokenExpiryInSeconds int32  `json:"access_token_expiry_in_seconds,omitempty"`
}
