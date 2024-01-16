package handler

import (
	"commune"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Hub *commune.Hub // websocket hub
	// TODO service
}

func (h *Handler) InitRoutes() *gin.Engine {
	app := gin.New()

	app.GET("/ws/:room", func(c *gin.Context) {
		room := c.Param("room")

		commune.ServeWs(h.Hub, c.Writer, c.Request, room)
	})

	return app
}
