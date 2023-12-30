package api

import (
	db "github.com/DMV-Nicolas/robotgram/backend/db/mongo"
	"github.com/DMV-Nicolas/robotgram/backend/token"
	"github.com/DMV-Nicolas/robotgram/backend/util"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config     util.Config
	queries    db.Querier
	tokenMaker token.Maker
	router     *echo.Echo
}

func NewServer(config util.Config, queries db.Querier) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{
		config:     config,
		queries:    queries,
		tokenMaker: tokenMaker,
	}

	e := echo.New()

	e.Validator = NewCustomValidator(validator.New())

	server.setupRouter(e)

	return server, nil
}

func (server *Server) setupRouter(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	v1.GET("/", server.Home)

	v1.POST("/users", server.CreateUser)
	v1.POST("/users/login", server.LoginUser)
	v1.GET("/users/:username", server.GetUser)
	v1.GET("/users", server.ListUsers)

	v1.POST("/posts", authMiddleware(server.CreatePost, server.tokenMaker))
	v1.GET("/posts", server.ListPosts)
	v1.GET("/posts/:id", server.GetPost)
	v1.PUT("/posts/:id", authMiddleware(server.UpdatePost, server.tokenMaker))
	v1.DELETE("/posts/:id", authMiddleware(server.DeletePost, server.tokenMaker))

	v1.POST("/likes", authMiddleware(server.ToggleLike, server.tokenMaker))
	v1.GET("/likes/:target_id", server.ListLikes)
	v1.GET("/likes/:target_id/count", server.CountLikes)

	v1.POST("/token/refresh", server.RefreshToken)

	server.router = e
}

func (server *Server) Start(address string) error {
	return server.router.Start(address)
}
