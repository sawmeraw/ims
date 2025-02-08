package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sawmeraw/ims/internal/services"
)

type Handler struct {
	services *services.Services
}

func New(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Home(c echo.Context) error {

	// c.Render(200, "home")

	return nil
}
