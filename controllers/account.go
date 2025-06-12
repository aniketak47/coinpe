package controllers

import (
	"coinpe/models"
	errorConst "coinpe/pkg/error"
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
		c.JSON(http.StatusInternalServerError, errResponse.Generate(
			errorConst.ErrorInternalError,
			"error in getting account",
			errorConst.EmptyInterface,
		))
		return
	}

	if existingAccount != nil && existingAccount.ID != 0 {
		c.JSON(http.StatusCreated, existingAccount)
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
		c.JSON(http.StatusInternalServerError, errResponse.Generate(
			errorConst.ErrorInternalError,
			"error in creating account",
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

	c.JSON(http.StatusCreated, account)

}
