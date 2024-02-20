package handler

import (
	"commune/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) Create(c *gin.Context) {
	var input entity.Message

	if err := c.BindJSON(&input); err != nil {
		//requestInput(body).invalidInput(c, err)
		logrus.Info(err)
		return
	}

	if _, err := h.services.Create(input); err != nil {
		//newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logrus.Info(err)
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) GetById(c *gin.Context) {
	id := entity.ObjectID(c.Param("ID"))

	msg, err := h.services.Get(id)
	if err != nil {
		logrus.Info(err)
		return
	}

	c.JSON(http.StatusOK, msg)
}

func (h *Handler) Get(c *gin.Context) {
	list := h.services.GetList()

	c.JSON(http.StatusOK, list)
}
