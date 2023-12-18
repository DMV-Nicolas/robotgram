package api

import (
	"context"
	"errors"
	"net/http"

	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type createPostRequest struct {
	Images      []string `json:"images"`
	Videos      []string `json:"videos" `
	Description string   `json:"description"`
}

func (server *Server) CreatePost(c echo.Context) error {
	req := new(createPostRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	userID, err := primitive.ObjectIDFromHex(c.Response().Header().Get(authorizationPayloadKey))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	arg := db.CreatePostParams{
		UserID:      userID,
		Images:      req.Images,
		Videos:      req.Videos,
		Description: req.Description,
	}

	result, err := server.queries.CreatePost(context.TODO(), arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

type getPostRequest struct {
	ID string `param:"id"`
}

func (server *Server) GetPost(c echo.Context) error {
	req := new(getPostRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
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
	Limit int `query:"limit" validate:"required,min=1"`
}

func (server *Server) ListPosts(c echo.Context) error {
	req := new(listPostsRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	posts, err := server.queries.ListPosts(context.TODO(), req.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, posts)
}

type updatePostRequest struct {
	ID          string   `json:"id" validate:"required"`
	Images      []string `json:"images"`
	Videos      []string `json:"videos" `
	Description string   `json:"description"`
}

func (server *Server) UpdatePost(c echo.Context) error {
	req := new(updatePostRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	gotPost, err := server.validPost(c, req.ID)
	if err != nil {
		return err
	}

	userID, err := primitive.ObjectIDFromHex(c.Response().Header().Get(authorizationPayloadKey))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if gotPost.UserID != userID {
		err = errors.New("account doesn't belong to the authenticated user")
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	arg := db.UpdatePostParams{
		ID:          gotPost.ID,
		Images:      req.Images,
		Videos:      req.Videos,
		Description: req.Description,
	}

	result, err := server.queries.UpdatePost(context.TODO(), arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

type deletePostRequest struct {
	ID string `json:"id" validate:"required"`
}

func (server *Server) DeletePost(c echo.Context) error {
	req := new(updatePostRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	gotPost, err := server.validPost(c, req.ID)
	if err != nil {
		return err
	}

	userID, err := primitive.ObjectIDFromHex(c.Response().Header().Get(authorizationPayloadKey))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if gotPost.UserID != userID {
		err = errors.New("account doesn't belong to the authenticated user")
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	result, err := server.queries.DeletePost(context.TODO(), gotPost.ID)
	if err != nil {
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
