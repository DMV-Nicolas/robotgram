package api

import (
	"context"
	"errors"
	"net/http"

	db "github.com/DMV-Nicolas/robotgram/backend/db/mongo"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

type updateCommentRequest struct {
	ID      string `param:"id" validate:"required,len=24"`
	Content string `json:"content" validate:"required"`
}

func (server *Server) UpdateComment(c echo.Context) error {
	req := new(updateCommentRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	gotComment, err := server.validComment(c, req.ID)
	if err != nil {
		return err
	}

	payload, err := getAuthorizationPayload(c)
	if err != nil {
		return err
	}

	if gotComment.UserID != payload.UserID {
		err = errors.New("account doesn't belong to the authenticated user")
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	arg := db.UpdateCommentParams{
		ID:      gotComment.ID,
		Content: req.Content,
	}

	result, err := server.queries.UpdateComment(context.TODO(), arg)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

type deleteCommentRequest struct {
	ID string `param:"id" validate:"required,len=24"`
}

func (server *Server) DeleteComment(c echo.Context) error {
	req := new(deleteCommentRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	gotComment, err := server.validComment(c, req.ID)
	if err != nil {
		return err
	}

	payload, err := getAuthorizationPayload(c)
	if err != nil {
		return err
	}

	if gotComment.UserID != payload.UserID {
		err = errors.New("account doesn't belong to the authenticated user")
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	result, err := server.queries.DeleteComment(context.TODO(), gotComment.ID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func (server *Server) validComment(c echo.Context, idStr string) (db.Comment, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		err = echo.NewHTTPError(http.StatusBadRequest, err)
		return db.Comment{}, err
	}

	comment, err := server.queries.GetComment(context.TODO(), id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = echo.NewHTTPError(http.StatusNotFound, err)
			return db.Comment{}, err
		}
		err = echo.NewHTTPError(http.StatusInternalServerError, err)
		return db.Comment{}, err
	}

	return comment, nil
}
