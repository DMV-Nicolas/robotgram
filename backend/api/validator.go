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
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return nil
}

// BindAndValidate bind and validate the given request
func bindAndValidate(c echo.Context, req interface{}) error {
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	return nil
}
