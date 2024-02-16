package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type refreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type refreshTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) RefreshToken(c echo.Context) error {
	req := new(refreshTokenRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	session, err := server.queries.GetSession(context.TODO(), refreshPayload.ID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if session.IsBlocked {
		err := errors.New("blocked session")
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	if session.UserID != refreshPayload.UserID {
		err := errors.New("incorrect session user")
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	if session.RefreshToken != req.RefreshToken {
		err := errors.New("mismatched session token")
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(session.UserID, server.config.AccessTokenDuration)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := refreshTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiresAt,
	}

	return c.JSON(http.StatusOK, res)
}

func (server *Server) GetTokenData(c echo.Context) error {
	payload, err := getAuthorizationPayload(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, payload)
}
