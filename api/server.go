package api

import (
	"net/http"

	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	"github.com/DMV-Nicolas/tinygram/util"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Server struct {
	config  util.Config
	queries db.Querier
	router  *echo.Echo
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func NewServer(config util.Config, queries db.Querier) (*Server, error) {
	server := &Server{
		config:  config,
		queries: queries,
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	server.setupRouter(e)

	return server, nil
}

func (server *Server) setupRouter(e *echo.Echo) {
	e.GET("/", server.Home)

	e.POST("/users", server.CreateUser)

	server.router = e
}

func (server *Server) Start(address string) error {
	return server.router.Start(address)
}
