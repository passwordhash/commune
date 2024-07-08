package handler

import (
	"commune/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) Create(c *gin.Context) {
	var input entity.MessageCreate

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	ID, exists := c.Get(idCtx)
	authorID, ok := ID.(entity.ObjectID)
	if !exists || !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user is unauthorized")
		return
	}
	if _, err := h.services.Create(input, authorID); err != nil {
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
	var list []entity.Message

	additionalQuery := c.Query("additional")
	isAdditional, _ := strconv.ParseBool(additionalQuery)

	if !isAdditional {
		list = h.services.GetList()
	}

	c.JSON(http.StatusOK, list)
}
