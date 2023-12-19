package api

import (
	"context"
	"net/http"

	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	gotPost, err := server.validPost(c, req.PostID)
	if err != nil {
		return err
	}

	arg := db.CreateLikeParams{
		UserID: payload.UserID,
		PostID: gotPost.ID,
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

type listLikesRequest struct {
	PostID string `query:"post_id" validate:"required,len=24"`
	Offset int64  `query:"offset" validate:"min=0"`
	Limit  int64  `query:"limit" validate:"min=1"`
}

func (server *Server) ListLikes(c echo.Context) error {
	req := new(listLikesRequest)
	if err := bindAndValidate(c, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	gotPost, err := server.validPost(c, req.PostID)
	if err != nil {
		return err
	}

	arg := db.ListLikesParams{
		PostID: gotPost.ID,
		Offset: req.Offset,
		Limit:  req.Limit,
	}

	posts, err := server.queries.ListLikes(context.TODO(), arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, posts)
}

type deleteLikeRequest struct {
	ID string `json:"id" validate:"required,len=24"`
}

func (server *Server) DeleteLike(c echo.Context) error {
	req := new(deleteLikeRequest)
	if err := bindAndValidate(c, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	payload, err := getAuthorizationPayload(c)
	if err != nil {
		return err
	}

	gotLike, err := server.validLike(c, req.ID)
	if err != nil {
		return err
	}

	if gotLike.UserID != payload.UserID {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	result, err := server.queries.DeleteLike(context.TODO(), gotLike.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func (server *Server) validLike(c echo.Context, idStr string) (db.Like, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		err = echo.NewHTTPError(http.StatusBadRequest, err)
		return db.Like{}, err
	}

	like, err := server.queries.GetLike(context.TODO(), id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = echo.NewHTTPError(http.StatusNotFound, err)
			return db.Like{}, err
		}
		err = echo.NewHTTPError(http.StatusInternalServerError, err)
		return db.Like{}, err
	}

	return like, nil
}