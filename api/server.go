package api

import (
	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	"github.com/DMV-Nicolas/tinygram/token"
	"github.com/DMV-Nicolas/tinygram/util"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:52330"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", server.Home)

	e.POST("/users", server.CreateUser)
	e.POST("/users/login", server.LoginUser)
	e.GET("/users/:username", server.GetUser)

	e.POST("/posts", authMiddleware(server.CreatePost, server.tokenMaker))
	e.GET("/posts", server.ListPosts)
	e.GET("/posts/:id", server.GetPost)

	server.router = e
}

func (server *Server) Start(address string) error {
	return server.router.Start(address)
}
