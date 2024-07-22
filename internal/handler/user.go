package handler

import (
	"commune/internal/entity"
	"commune/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signUpResponse struct {
	Token entity.JWTToken     `json:"token"`
	User  entity.UserResponse `json:"user"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var input entity.UserCreate

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, passcode, err := h.services.User.SignUp(input)
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

	// TODO: делать это в горутине (нам не особо важно, чтобы письмо
	//       точно отправилось. Пользователь в любой момоент может
	//       запросить отправку письма в любой другой момент)
	h.services.Email.SendCode(input.Email, passcode)

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

	authData, err := h.services.User.Authenticate(input)
	if errors.Is(err, service.UserNotFound) {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, signUpResponse{
		Token: authData.Token,
		User:  authData.User,
	})
}

func (h *Handler) ResetPasscode(c *gin.Context) {
	var input entity.UserEmail

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	newPasscode, err := h.services.User.UpdatePasscode(input)
	if errors.Is(err, service.UserNotFound) {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: делать это в горутине (нам не особо важно, чтобы письмо
	//       точно отправилось. Пользователь в любой момоент может
	//       запросить отправку письма в любой другой момент)
	h.services.Email.SendCode(input.Email, newPasscode)

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) Check(c *gin.Context) {
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
