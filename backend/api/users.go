package api

import (
	"context"
	"net/http"
	"time"

	db "github.com/DMV-Nicolas/robotgram/backend/db/mongo"
	"github.com/DMV-Nicolas/robotgram/backend/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
		Avatar:         req.Avatar,
		Gender:         req.Gender,
	}

	result, err := server.queries.CreateUser(context.TODO(), arg)
	if err != nil {
		if err == db.ErrUsernameTaken || err == db.ErrEmailTaken {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

type loginUserRequest struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
}

type loginUserResponse struct {
	SessionID             string    `json:"session_id"`
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}

func (server *Server) LoginUser(c echo.Context) error {
	req := new(loginUserRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	addr, isMail := util.ValidMailAddress(req.UsernameOrEmail)

	var err error
	var user db.User
	if isMail {
		user, err = server.queries.GetUser(context.TODO(), "email", addr)
	} else {
		user, err = server.queries.GetUser(context.TODO(), "username", req.UsernameOrEmail)
	}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := util.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(user.ID, server.config.AccessTokenDuration)
	if err != nil {
		// impossible
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(user.ID, server.config.RefreshTokenDuration)
	if err != nil {
		// impossible
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	arg := db.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserID:       refreshPayload.UserID,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIP:     "",
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiresAt,
	}

	_, err = server.queries.CreateSession(context.TODO(), arg)

	res := loginUserResponse{
		SessionID:             refreshPayload.ID.Hex(),
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiresAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiresAt,
	}

	return c.JSON(http.StatusOK, res)
}

type getUserRequest struct {
	ID string `param:"id" validate:"required,len=24"`
}

type getUserResponse struct {
	ID          primitive.ObjectID `json:"id"`
	Username    string             `json:"username"`
	FullName    string             `json:"full_name"`
	Email       string             `json:"email"`
	Avatar      string             `json:"avatar"`
	Description string             `json:"description"`
	Gender      string             `json:"gender"`
	CreatedAt   time.Time          `json:"created_at"`
}

func (server *Server) GetUser(c echo.Context) error {
	req := new(getUserRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := server.queries.GetUser(context.TODO(), "_id", id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := getUserResponse{
		ID:          user.ID,
		Username:    user.Username,
		FullName:    user.FullName,
		Email:       user.Email,
		Avatar:      user.Avatar,
		Description: user.Description,
		Gender:      user.Gender,
		CreatedAt:   user.CreatedAt,
	}

	return c.JSON(http.StatusOK, res)
}

type listUsersRequest struct {
	Offset int64 `query:"offset" validate:"min=0"`
	Limit  int64 `query:"limit" validate:"min=1"`
}

func (server *Server) ListUsers(c echo.Context) error {
	req := new(listUsersRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	arg := db.ListUsersParams{
		Offset: req.Offset,
		Limit:  req.Limit,
	}

	users, err := server.queries.ListUsers(context.TODO(), arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}
