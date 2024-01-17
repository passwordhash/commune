package handler

import (
	"commune/internal/service"
	"commune/internal/websocket"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	hub      *websocket.Hub // websocket hub
	services *service.Service
}

func NewHandler(hub *websocket.Hub, services *service.Service) *Handler {
	return &Handler{hub: hub, services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	app := gin.New()

	app.GET("/ws/:room", func(c *gin.Context) {
		room := c.Param("room")

		websocket.ServeWs(h.hub, c.Writer, c.Request, room)
	})

	return app
}
