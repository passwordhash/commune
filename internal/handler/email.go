package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) TestEmail(c *gin.Context) {
	email := c.Param("email")

	err := h.services.Email.SendCode(email, "f5B2Gg")
	if err != nil {
		logrus.Info(err)
		c.JSON(400, map[string]string{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"email": email,
	})
}
