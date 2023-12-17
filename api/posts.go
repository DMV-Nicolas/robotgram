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
	UserID      string   `json:"user_id" validate:"required"`
	Images      []string `json:"images"`
	Videos      []string `json:"videos" `
	Description string   `json:"description"`
}

func (server *Server) CreatePost(c echo.Context) error {
	req := new(createPostRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userID, err := primitive.ObjectIDFromHex(req.UserID)
	_, err = server.queries.GetUser(context.TODO(), "_id", userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	arg := db.CreatePostParams{
		UserID:      userID,
		Images:      req.Images,
		Videos:      req.Videos,
		Description: req.Description,
	}

	result, err := server.queries.CreatePost(context.TODO(), arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}
