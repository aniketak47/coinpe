package controllers

import (
	"coinpe/pkg/jwtauth"
	"coinpe/pkg/logger"
	"coinpe/pkg/utils"
	"time"
)

// createToken: Creates the respective token and persists in redis
func (b *BaseController) createToken(request *CreateTokenRequest) (*CreateTokenResponse, error) {
	var (
		expiryTime      = time.Now()
		expiryInSeconds int
	)

	if request.ExpiryTimeInSeconds != nil {
		expiryInSeconds = int(*request.ExpiryTimeInSeconds)
		expiryTime = expiryTime.Add(time.Second * time.Duration(*request.ExpiryTimeInSeconds))
	} else if request.IsPartial {
		expiryInSeconds = b.Config.JWTConfiguration.PartialAuthAccessTokenExpiryInSeconds
		expiryTime = expiryTime.Add(time.Second * time.Duration(expiryInSeconds))
	} else if !request.IsPartial && request.TokenType == TokenTypeAccess {
		expiryInSeconds = b.Config.JWTConfiguration.FullAuthAccessTokenExpiryInSeconds
		expiryTime = expiryTime.Add(time.Second * time.Duration(expiryInSeconds))
	} else if !request.IsPartial && request.TokenType == TokenTypeRefresh {
		expiryInSeconds = b.Config.JWTConfiguration.FullAuthRefreshTokenExpiryInSeconds
		expiryTime = expiryTime.Add(time.Second * time.Duration(expiryInSeconds))

	}

	token, err := jwtauth.NewTokenWithClaims([]byte(b.Config.JWTConfiguration.SecretKey),
		*request.CustomClaims, expiryTime)
	if err != nil {
		logger.Error("unable to create token with claims ", err)
		return nil, err
	}

	return &CreateTokenResponse{
		Token:           utils.String(token),
		ExpiryInSeconds: expiryInSeconds,
	}, nil

}
