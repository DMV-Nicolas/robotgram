package api

import (
	"context"
	"net/http"
	"time"

	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	"github.com/DMV-Nicolas/tinygram/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type createUserRequest struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,min=8"`
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Avatar   string `json:"avatar" validate:"required"`
	Gender   string `json:"gender" validate:"required,oneof=male female"`
}

func (server *Server) CreateUser(c echo.Context) error {
	req := new(createUserRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
		Avatar:         req.Avatar,
		Gender:         req.Gender,
	}

	result, err := server.querier.CreateUser(context.TODO(), arg)
	if err != nil {
		if err == db.ErrUsernameTaken || err == db.ErrEmailTaken {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

type loginUserRequest struct {
	Username string `json:"username" validate:"required_without=Email"`
	Email    string `json:"email" validate:"required_without=Username"`
	Password string `json:"password" validate:"required,min=8"`
}

type loginUserResponse struct {
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"acess_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}

func (server *Server) LoginUser(c echo.Context) error {
	req := new(loginUserRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var err error
	var user db.User
	if req.Username != "" {
		user, err = server.querier.GetUser(context.TODO(), "username", req.Username)
	} else {
		user, err = server.querier.GetUser(context.TODO(), "email", req.Email)
	}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := util.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		// imposible
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.RefreshTokenDuration)
	if err != nil {
		// imposible
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := loginUserResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiresAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiresAt,
	}

	return c.JSON(http.StatusOK, res)
}
