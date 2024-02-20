package handler

import (
	"commune/internal/entity"
	"commune/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signUpResponse struct {
	Token entity.JWTToken `json:"token"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var input entity.UserCreate

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.services.User.SignUp(input)
	if errors.Is(err, service.UserAlreadyExists) {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}
	if errors.Is(err, service.UserCreationError) {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//logrus.Info(token)

	c.JSON(http.StatusOK, signUpResponse{
		Token: token,
	})
}

func (h *Handler) LogIn(c *gin.Context) {
	var input entity.UserAuth

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.services.User.Authenticate(input)
	if errors.Is(err, service.UserNotFound) {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, signUpResponse{
		Token: token,
	})
}
