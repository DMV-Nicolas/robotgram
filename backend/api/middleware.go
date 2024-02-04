package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/DMV-Nicolas/robotgram/backend/token"
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
			err := errors.New("unsupported authorization type: " + authorizationType)
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		accessToken := fields[1]
		fmt.Printf("Access Token: \"%s\"\n", accessToken)
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		payloadJSON, err := json.Marshal(payload)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		c.Response().Header().Set(authorizationPayloadKey, string(payloadJSON))

		return next(c)
	}
}

func getAuthorizationPayload(c echo.Context) (*token.Payload, error) {
	payloadJSON := c.Response().Header().Get(authorizationPayloadKey)
	payload := new(token.Payload)
	err := json.Unmarshal([]byte(payloadJSON), payload)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return payload, nil
}
