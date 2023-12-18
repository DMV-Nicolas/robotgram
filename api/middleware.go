package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/DMV-Nicolas/tinygram/token"
	"github.com/labstack/echo/v4"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(next echo.HandlerFunc, tokenMaker token.Maker) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get(authorizationHeaderKey)
		if authHeader == "" {
			err := errors.New("authorization header not provided")
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := errors.New("unsopported authorization type: " + authorizationType)
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		c.Response().Header().Set(authorizationPayloadKey, payload.UserID.Hex())

		return next(c)
	}
}
