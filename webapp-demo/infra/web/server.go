package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Router *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return &Server{Router: e}
}

func (s *Server) ListenAndServe() {
	e := s.Router
	e.Logger.Fatal(e.Start(":8080"))
}
