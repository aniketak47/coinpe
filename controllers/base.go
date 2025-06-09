package controllers

import (
	"coinpe/pkg/config"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type BaseController struct {
	DB         *gorm.DB
	Config     config.Config
	Validator  *validator.Validate
	Translator *ut.Translator
}
