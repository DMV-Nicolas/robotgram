package api

import (
	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	"github.com/DMV-Nicolas/tinygram/util"
	"github.com/labstack/echo/v4"
)

type Server struct {
	config  util.Config
	queries db.Querier
	router  *echo.Echo
}

func NewServer(config util.Config, queries db.Querier) (*Server, error) {
	server := &Server{
		config:  config,
		queries: queries,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	e := echo.New()

	e.GET("/", server.Home)

	server.router = e
}

func (server *Server) Start(address string) error {
	return server.router.Start(address)
}
