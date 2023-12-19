package api

import (
	"context"
	"net/http"

	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createLikeRequest struct {
	PostID string `json:"post_id" validate:"required,len=24"`
}

func (server *Server) CreateLike(c echo.Context) error {
	req := new(createLikeRequest)
	if err := bindAndValidate(c, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	payload, err := getAuthorizationPayload(c)
	if err != nil {
		return err
	}

	postID, err := primitive.ObjectIDFromHex(req.PostID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	arg := db.CreateLikeParams{
		UserID: payload.UserID,
		PostID: postID,
	}

	result, err := server.queries.CreateLike(context.TODO(), arg)
	if err != nil {
		if err == db.ErrDuplicatedLike {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}
