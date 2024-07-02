package handler

import (
	"commune/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) Create(c *gin.Context) {
	var input entity.MessageCreate
	// TODO: проверка/сравнение input.authorID с UserId из контекса

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if _, err := h.services.Create(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
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
