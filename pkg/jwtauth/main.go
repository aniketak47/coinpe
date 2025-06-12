package jwtauth

import (
	"coinpe/pkg/logger"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewTokenWithClaims(secretKey []byte, customClaims CustomClaims,
	expires time.Time) (*string, error) {
	claims := JWTTokenClaims{
		customClaims,
		jwt.RegisteredClaims{
			Issuer:    JWTIssuer,
			ExpiresAt: jwt.NewNumericDate(expires),
		},
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Signed string
	signedString, err := token.SignedString(secretKey)
	if err != nil {
		logger.Error("error in creating new jwt token ", err)
		return nil, err
	}
	return &signedString, nil
}

func ParseToken(token string, secretKey []byte, ignoreValidity ...bool) (*CustomClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &JWTTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		logger.Error("unable to parse token ", err)
		if err.Error() == "token has invalid claims: token is expired" || err.Error() == jwt.ErrTokenExpired.Error() {
			if len(ignoreValidity) > 0 && ignoreValidity[0] {
				logger.Info("overriding expiration check")
				claims, ok := parsedToken.Claims.(*JWTTokenClaims)
				if ok {
					return &claims.CustomClaims, nil
				}
			}
		}
		return nil, err
	}
	if claims, ok := parsedToken.Claims.(*JWTTokenClaims); ok {
		if parsedToken.Valid {
			return &claims.CustomClaims, nil
		} else if !parsedToken.Valid && len(ignoreValidity) > 0 && ignoreValidity[0] {
			logger.Info("overriding validity check")
			return &claims.CustomClaims, nil
		}
	}
	return nil, errors.New("invalid token")
}
