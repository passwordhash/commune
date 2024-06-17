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

	app.Use(CORSMiddleware())

	app.GET("/chat", func(c *gin.Context) {
		c.File("./index.html")
	})

	app.GET("/ws/:room", func(c *gin.Context) {
		room := c.Param("room")

		websocket.ServeWs(h.hub, *h.services, c.Writer, c.Request, room)
	})

	auth := app.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.LogIn)
	}

	api := app.Group("/api")
	{
		api.POST("/new", h.Create)
		api.GET("/message/:ID", h.GetById)
		api.GET("/list", h.Get)

	}

	return app
}
