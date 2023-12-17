package api

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CustomValidator is a custom validator
type CustomValidator struct {
	validator *validator.Validate
}

// NewCustomValidator creates a new CustomValidator
func NewCustomValidator(validator *validator.Validate) *CustomValidator {
	return &CustomValidator{
		validator: validator,
	}
}

// Validate validates the struct data
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
