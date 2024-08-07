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
		auth.POST("/reset", h.ResetPasscode)
		auth.GET("/check", h.userIdentity, h.Check)
	}

	api := app.Group("/api", h.userIdentity)
	{
		message := api.Group("/message")
		{
			message.POST("/new", h.Create)
			message.GET("/list", h.Get)
			message.GET("/:ID", h.GetById)
		}

	}

	return app
}
