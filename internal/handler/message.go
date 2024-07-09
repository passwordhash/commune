package handler

import (
	"commune/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

const withAdditionsQueryP = "withAdditions"

func (h *Handler) Create(c *gin.Context) {
	var input entity.MessageCreate
	var output entity.Message

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
	output, err := h.services.Create(input, authorID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//c.JSON(http.StatusOK, statusResponse{"ok"})
	c.JSON(http.StatusOK, output)
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
	var err error

	additionalQuery := c.Query(withAdditionsQueryP)
	isAdditional, _ := strconv.ParseBool(additionalQuery)

	if isAdditional {
		list, err = h.services.GetListWithAdditions()
	} else {
		list = h.services.GetList()
	}

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}
