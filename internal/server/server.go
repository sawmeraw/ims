package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sawmeraw/ims/internal/config"
	"github.com/sawmeraw/ims/internal/database"
	"github.com/sawmeraw/ims/internal/handlers"
	"github.com/sawmeraw/ims/internal/middleware"
	"github.com/sawmeraw/ims/internal/services"
)

type Server struct {
	cfg      *config.Config
	db       *database.Database
	echo     *echo.Echo
	services *services.Services
}

// load the config, initialize the database connection, init services and start the echo server
func Init() (*Server, error) {
	cfg, err := config.Load()

	if err != nil {
		return nil, err
	}

	db, err := database.New(cfg)
	if err != nil {
		return nil, err
	}
	log.Println("Database connection successful")

	e := echo.New()

	services := services.NewServices(db)

	return &Server{
		cfg:      cfg,
		db:       db,
		echo:     e,
		services: services,
	}, nil

}

// returning a server pointer to chain calls, just to not confuse myself
func (s *Server) SetupRoutes() *Server {
	// s.echo.Static("static")
	h := handlers.New(s.services)
	s.echo.GET("/", h.Home)
	return s
}

func (s *Server) SetupMiddleware() *Server {
	middleware.Setup(s.echo)
	return s
}

func (s *Server) Start() error {
	return s.echo.Start(":" + s.cfg.Port)
}
