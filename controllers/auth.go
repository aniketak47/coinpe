package controllers

import (
	"coinpe/models"
	errorConst "coinpe/pkg/error"
	otphelpers "coinpe/pkg/helpers/otp_helpers"
	"coinpe/pkg/jwtauth"
	"coinpe/pkg/logger"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (b *BaseController) Authenticate(c *gin.Context) {
	var (
		email        string
		phone        string
		refreshToken string
		request      = AuthenticateRequest{}
		errResponse  = errorConst.ErrorResponse{}
		accountRepo  = models.InitAccountRepo(b.DB)
	)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		logger.Error("unable to bind request | err: ", err)
		c.JSON(http.StatusBadRequest,
			errResponse.Generate(
				errorConst.ErrorBindingRequest,
				errorConst.ErrorText(errorConst.ErrorBindingRequest),
				errorConst.EmptyInterface))
		return
	}

	if strings.Contains(request.Username, "@") {
		email = request.Username
	} else {
		phone = request.Username
	}

	if !slices.Contains([]string{string(models.RoleTypeSuperAdmin), string(models.RoleTypeAdmin), string(models.RoleTypeCustomer)}, request.Role) {
		logger.Error("invalid role")
		c.JSON(http.StatusBadRequest, errResponse.Generate(
			errorConst.ErrorBadRequest,
			"invalid role",
			errorConst.EmptyInterface,
		))
		return
	}

	tx := b.DB.Begin()

	// check account exists or not
	account, err := accountRepo.FindOne(tx, email, phone, "")
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error("error in getting account | err: ", err)
		c.JSON(http.StatusInternalServerError, errResponse.Generate(
			errorConst.ErrorInternalError,
			"error in getting account",
			errorConst.EmptyInterface,
		))
		return
	}

	if err == gorm.ErrRecordNotFound || account.ID == 0 {
		logger.Error("account not found", err)
		c.JSON(http.StatusNotFound, AuthenticateResponse{
			GoTo: GotToCreateAccount,
		})
		return
	}

	accessTokenExpiryTime := int32(b.Config.JWTConfiguration.PartialAuthAccessTokenExpiryInSeconds)
	customClaims := &jwtauth.CustomClaims{
		Role:        request.Role,
		AccountUUID: account.UUID,
		Email:       email,
		PhoneNumber: phone,
		IsPartial:   true,
	}

	//create token
	accessToken, err := b.createToken(&CreateTokenRequest{
		TokenType:           TokenTypeAccess,
		ExpiryTimeInSeconds: &accessTokenExpiryTime,
		CustomClaims:        customClaims,
	})
	if err != nil {
		logger.Error("unable to create access token ", err)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, errResponse.Generate(
			errorConst.ErrorInternalError,
			"error in creating access token",
			errorConst.EmptyInterface,
		))
		return
	}

	err = tx.Commit().Error
	if err != nil {
		logger.Error("unable to commit | err: ", err)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, errResponse.Generate(
			errorConst.ErrorInternalError,
			"error in commiting",
			errorConst.EmptyInterface,
		))
		return
	}

	c.JSON(http.StatusOK, AuthenticateResponse{
		AccessToken:                accessToken.Token,
		RefreshToken:               refreshToken,
		AccessTokenExpiryInSeconds: int32(accessToken.ExpiryInSeconds),
		VerificationChannel:        "sms",
		Handle:                     request.Username,
		GoTo:                       GoToVerifyAccount,
	})

}

func (b *BaseController) VerifyAuthenticate(c *gin.Context) {
	var (
		request         = VerifyAuthRequest{}
		errResponse     = errorConst.ErrorResponse{}
		accountRepo     = models.InitAccountRepo(b.DB)
		credentialsRepo = models.InitCredentialRepo(b.DB)
	)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		logger.Error("error in binding request | err: ", err)
		c.JSON(http.StatusBadRequest,
			errResponse.Generate(
				errorConst.ErrorBindingRequest,
				errorConst.ErrorText(errorConst.ErrorBindingRequest),
				errorConst.EmptyInterface))
		return
	}

	// Parse token
	parsedTokenClaims, err := jwtauth.ParseToken(request.AccessToken, []byte(b.Config.JWTConfiguration.SecretKey))
	if err != nil {
		logger.Error("error in getting parsed token | err: ", err)
		c.JSON(http.StatusInternalServerError, errResponse.Generate(
			errorConst.ErrorInternalError,
			"error in getting parsed token",
			errorConst.EmptyInterface,
		))
		return
	}

	account, err := accountRepo.Get(&models.Account{
		UUID: parsedTokenClaims.AccountUUID,
	})
	if err != nil {
		logger.Error("unable to get account | err: ", err)
		c.JSON(http.StatusInternalServerError, errResponse.Generate(
			errorConst.ErrorInternalError,
			"error in getting account",
			errorConst.EmptyInterface,
		))
		return
	}

	// Get credential
	credentials, err := credentialsRepo.Get(&models.Credential{
		AccountID: &account.ID,
		Type:      models.CredentialsTypeOTPSecret,
	})
	if err != nil {
		logger.Error("error in creating otp secret | err: ", err)
		c.JSON(http.StatusInternalServerError, errResponse.Generate(
			errorConst.ErrorInternalError,
			"error in creating otp secret",
			errorConst.EmptyInterface,
		))
		return
	}

	secretKey := credentials.Password
	accountUUID := account.UUID

	isOTPValid, err := otphelpers.ValidateOTP(secretKey, request.Otp, b.Config.ShouldMock())
	if err != nil {
		logger.Error("unable to validate user otp | err: ", err)
		c.JSON(http.StatusInternalServerError, errResponse.Generate(
			errorConst.ErrorInternalError,
			"unable to validate user otp",
			errorConst.EmptyInterface,
		))
		return
	}

	if !isOTPValid {
		logger.Info("invalid otp provided")
		c.JSON(http.StatusInternalServerError, errResponse.Generate(
			errorConst.ErrorInternalError,
			"invalid otp provided",
			errorConst.EmptyInterface,
		))
		return
	}

	accessTokenRequest := CreateTokenRequest{
		TokenType: TokenTypeAccess,
		CustomClaims: &jwtauth.CustomClaims{
			Role:        string(parsedTokenClaims.Role),
			AccountUUID: accountUUID,
		},
	}

	if request.AccessTokenExpiryInSeconds > 0 {
		expiryTime := request.AccessTokenExpiryInSeconds
		accessTokenRequest.ExpiryTimeInSeconds = &expiryTime
	}

	// Generate a full scoped access and refresh token
	accessTokenResponse, err := b.createToken(&accessTokenRequest)
	if err != nil {
		logger.Error("unable to create access token | err: ", err)
		c.JSON(http.StatusInternalServerError, errResponse.Generate(
			errorConst.ErrorInternalError,
			"unable to create access token",
			errorConst.EmptyInterface,
		))
		return
	}

	c.JSON(http.StatusOK, AuthenticateResponse{
		AccessToken:                accessTokenResponse.Token,
		AccessTokenExpiryInSeconds: int32(accessTokenResponse.ExpiryInSeconds),
		GoTo:                       GoToContinue,
	})
}
