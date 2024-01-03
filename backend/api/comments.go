package api

import (
	"context"
	"net/http"

	db "github.com/DMV-Nicolas/robotgram/backend/db/mongo"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createCommentRequest struct {
	TargetID string `json:"target_id" validate:"required,len=24"`
	Content  string `json:"content" validate:"required"`
}

func (server *Server) CreateComment(c echo.Context) error {
	req := new(createCommentRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	payload, err := getAuthorizationPayload(c)
	if err != nil {
		return err
	}

	targetID, err := primitive.ObjectIDFromHex(req.TargetID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	arg := db.CreateCommentParams{
		UserID:   payload.UserID,
		TargetID: targetID,
		Content:  req.Content,
	}

	result, err := server.queries.CreateComment(context.TODO(), arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

type listCommentsRequest struct {
	TargetID string `param:"target_id" validate:"required,len=24"`
	Offset   int64  `query:"offset" validate:"min=0"`
	Limit    int64  `query:"limit" validate:"min=1"`
}

func (server *Server) ListComments(c echo.Context) error {
	req := new(listCommentsRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	targetID, err := primitive.ObjectIDFromHex(req.TargetID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	arg := db.ListCommentsParams{
		TargetID: targetID,
		Offset:   req.Offset,
		Limit:    req.Limit,
	}

	comments, err := server.queries.ListComments(context.TODO(), arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, comments)
}
