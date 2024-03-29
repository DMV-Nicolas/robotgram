package api

import (
	"context"
	"net/http"

	db "github.com/DMV-Nicolas/robotgram/backend/db/mongo"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type toggleLikeRequest struct {
	TargetID string `json:"target_id" validate:"required,len=24"`
}

type toggleLikeResponse struct {
	CreatedResult *mongo.InsertOneResult
	DeletedResult *mongo.DeleteResult
}

func (server *Server) ToggleLike(c echo.Context) error {
	req := new(toggleLikeRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	targetID, err := primitive.ObjectIDFromHex(req.TargetID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	payload, err := getAuthorizationPayload(c)
	if err != nil {
		return err
	}

	arg := db.ToggleLikeParams{
		UserID:   payload.UserID,
		TargetID: targetID,
	}

	createdResult, deletedResult, err := server.queries.ToggleLike(context.TODO(), arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := toggleLikeResponse{
		CreatedResult: createdResult,
		DeletedResult: deletedResult,
	}

	return c.JSON(http.StatusOK, res)
}

type listLikesRequest struct {
	TargetID string `param:"target_id" validate:"required,len=24"`
	Offset   int64  `query:"offset" validate:"min=0"`
	Limit    int64  `query:"limit" validate:"min=1"`
}

func (server *Server) ListLikes(c echo.Context) error {
	req := new(listLikesRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	targetID, err := primitive.ObjectIDFromHex(req.TargetID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	arg := db.ListLikesParams{
		TargetID: targetID,
		Offset:   req.Offset,
		Limit:    req.Limit,
	}

	posts, err := server.queries.ListLikes(context.TODO(), arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, posts)
}

type countLikesRequest struct {
	TargetID string `param:"target_id" validate:"required,len=24"`
}

func (server *Server) CountLikes(c echo.Context) error {
	req := new(countLikesRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	targetID, err := primitive.ObjectIDFromHex(req.TargetID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	nLikes, err := server.queries.CountLikes(context.TODO(), targetID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nLikes)
}

type isLikedRequest struct {
	TargetID string `param:"target_id" validate:"required,len=24"`
}

func (server *Server) IsLiked(c echo.Context) error {
	req := new(isLikedRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	targetID, err := primitive.ObjectIDFromHex(req.TargetID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	payload, err := getAuthorizationPayload(c)
	if err != nil {
		return err
	}

	arg := db.IsLikedParams{
		UserID:   payload.UserID,
		TargetID: targetID,
	}

	_, liked, err := server.queries.IsLiked(context.TODO(), arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, liked)
}
