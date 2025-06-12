package controllers

import (
	"coinpe/models"
	errorConst "coinpe/pkg/error"
	"coinpe/pkg/jwtauth"
	"coinpe/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (b *BaseController) CreateAccount(c *gin.Context) {
	var (
		roleID      uint64
		request     = CreateAccountRequest{}
		errResponse = errorConst.ErrorResponse{}
		accountRepo = models.InitAccountRepo(b.DB)
		walletRepo  = models.InitWalletRepo(b.DB)
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

	switch request.Role {
	case string(models.RoleTypeSuperAdmin):
		roleID = uint64(models.RoleSuperAdmin)

	case string(models.RoleTypeAdmin):
		roleID = uint64(models.RoleAdmin)

	case string(models.RoleTypeCustomer):
		roleID = uint64(models.RoleCustomer)

	default:
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
	existingAccount, err := accountRepo.FindOne(tx, request.Email, request.PhoneNumber, "")
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error("error in getting account | err: ", err)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, errResponse.Generate(
			errorConst.ErrorInternalError,
			"error in getting account",
			errorConst.EmptyInterface,
		))
		return
	}

	accessTokenExpiryTime := int32(b.Config.JWTConfiguration.PartialAuthAccessTokenExpiryInSeconds)

	if existingAccount != nil && existingAccount.ID != 0 {
		customClaims := &jwtauth.CustomClaims{
			Role:        request.Role,
			AccountUUID: existingAccount.UUID,
			Email:       existingAccount.Email,
			PhoneNumber: *existingAccount.PhoneNumber,
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

		c.JSON(http.StatusOK, AuthenticateResponse{
			AccessToken:                accessToken.Token,
			AccessTokenExpiryInSeconds: int32(accessToken.ExpiryInSeconds),
			VerificationChannel:        "sms",
			Handle:                     *existingAccount.PhoneNumber,
			GoTo:                       GoToVerifyAccount,
		})
		return
	}

	account := models.Account{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		PhoneNumber: &request.PhoneNumber,
		Email:       request.Email,
		RoleID:      roleID,
	}

	err = accountRepo.CreateWithTx(tx, &account)
	if err != nil {
		logger.Error("error in creating account | err: ", err)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, errResponse.Generate(
			errorConst.ErrorInternalError,
			"error in creating account",
			errorConst.EmptyInterface,
		))
		return
	}

	// create wallet
	err = walletRepo.CreateWithTx(tx, &models.Wallet{
		UserUUID:              account.UUID,
		Currency:              models.EntityINR,
		OverdraftLimitInCents: 10000, // initially giving ₹100 as overdraft
	})

	if err != nil {
		logger.Error("error in creating wallet | err: ", err)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, errResponse.Generate(
			errorConst.ErrorInternalError,
			"error in creating wallet",
			errorConst.EmptyInterface,
		))
		return
	}

	customClaims := &jwtauth.CustomClaims{
		Role:        request.Role,
		AccountUUID: account.UUID,
		Email:       account.Email,
		PhoneNumber: *account.PhoneNumber,
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
		logger.Error("error in commiting | err: ", err)
		c.JSON(http.StatusInternalServerError, errResponse.Generate(
			errorConst.ErrorInternalError,
			"error in getting account",
			errorConst.EmptyInterface,
		))
		return
	}

	c.JSON(http.StatusOK, AuthenticateResponse{
		AccessToken:                accessToken.Token,
		AccessTokenExpiryInSeconds: int32(accessToken.ExpiryInSeconds),
		VerificationChannel:        "sms",
		Handle:                     *existingAccount.PhoneNumber,
		GoTo:                       GoToVerifyAccount,
	})

}
