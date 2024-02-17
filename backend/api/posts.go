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

type createPostRequest struct {
	Images      []string `json:"images"`
	Description string   `json:"description"`
}

func (server *Server) CreatePost(c echo.Context) error {
	req := new(createPostRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	payload, err := getAuthorizationPayload(c)
	if err != nil {
		return err
	}

	arg := db.CreatePostParams{
		UserID:      payload.UserID,
		Images:      req.Images,
		Description: req.Description,
	}

	result, err := server.queries.CreatePost(context.TODO(), arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

type getPostRequest struct {
	ID string `param:"id" validate:"required,len=24"`
}

func (server *Server) GetPost(c echo.Context) error {
	req := new(getPostRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	post, err := server.queries.GetPost(context.TODO(), "_id", id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, post)
}

type listPostsRequest struct {
	Offset int64  `query:"offset" validate:"min=0"`
	Limit  int64  `query:"limit" validate:"min=1"`
	UserID string `query:"user_id"`
}

func (server *Server) ListPosts(c echo.Context) error {
	req := new(listPostsRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	var userID primitive.ObjectID
	var err error
	if req.UserID != "" && req.UserID != "<nil>" {
		userID, err = primitive.ObjectIDFromHex(req.UserID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
	}

	arg := db.ListPostsParams{
		Offset: req.Offset,
		Limit:  req.Limit,
		UserID: userID,
	}

	posts, err := server.queries.ListPosts(context.TODO(), arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, posts)
}

type updatePostRequest struct {
	ID          string   `param:"id" validate:"required,len=24"`
	Images      []string `json:"images"`
	Description string   `json:"description"`
}

func (server *Server) UpdatePost(c echo.Context) error {
	req := new(updatePostRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	gotPost, err := server.validPost(c, req.ID)
	if err != nil {
		return err
	}

	payload, err := getAuthorizationPayload(c)
	if err != nil {
		return err
	}

	if gotPost.UserID != payload.UserID {
		err = errors.New("account doesn't belong to the authenticated user")
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	arg := db.UpdatePostParams{
		ID:          gotPost.ID,
		Images:      req.Images,
		Description: req.Description,
	}

	result, err := server.queries.UpdatePost(context.TODO(), arg)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

type deletePostRequest struct {
	ID string `param:"id" validate:"required,len=24"`
}

func (server *Server) DeletePost(c echo.Context) error {
	req := new(deletePostRequest)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	gotPost, err := server.validPost(c, req.ID)
	if err != nil {
		return err
	}

	payload, err := getAuthorizationPayload(c)
	if err != nil {
		return err
	}

	if gotPost.UserID != payload.UserID {
		err = errors.New("account doesn't belong to the authenticated user")
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	result, err := server.queries.DeletePost(context.TODO(), gotPost.ID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func (server *Server) validPost(c echo.Context, idStr string) (db.Post, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		err = echo.NewHTTPError(http.StatusBadRequest, err)
		return db.Post{}, err
	}

	post, err := server.queries.GetPost(context.TODO(), "_id", id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = echo.NewHTTPError(http.StatusNotFound, err)
			return db.Post{}, err
		}
		err = echo.NewHTTPError(http.StatusInternalServerError, err)
		return db.Post{}, err
	}

	return post, nil
}
