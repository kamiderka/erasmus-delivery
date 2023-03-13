package dtos

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Register Consumer POST DTO
type RegisterConsumerDTO struct {
	FirstName   string 	`json:"firstName" validate:"required"`
	LastName	string  `json:"lastName" validate:"required"`
	Email		string	`json:"email" validate:"required,email"`
	Address		string	`json:"address"`
	City		string	`json:"city"`
	Password 	string 	`json:"password" validate:"required"`
}

// Login Consumer POST DTO
type LoginConsumerDTO struct {
	Email		string	`json:"email" validate:"required"`
	Password 	string 	`json:"password" validate:"required"`
}


func ValidateDTO(dto interface{}) ([]error) {
	validate := validator.New()
	errs_cst := []error{}

	err_org := validate.Struct(dto)
	if err_org != nil {
		for _, err := range err_org.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				errs_cst = append(errs_cst, errors.New(fmt.Sprintf("%s is required", err.Field())))
			case "email":
				errs_cst = append(errs_cst, errors.New(fmt.Sprintf("%s is not a valid email address", err.Field())))
			case "min":
				errs_cst = append(errs_cst, errors.New(fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())))
			}
		}
		return errs_cst
	}

	return nil
}

