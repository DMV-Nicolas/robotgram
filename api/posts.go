package api

import (
	"context"
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
