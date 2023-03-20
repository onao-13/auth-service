package controller

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"security-service/internals/app/entity/request"
	"security-service/internals/app/handler"
	user2 "security-service/internals/app/service"
)

// AuthController Контроллер для аутентификации
type AuthController struct {
	service *user2.AuthService
}

var log = logrus.New()

// NewAuthController Создание нового контроллера
func NewAuthController(service *user2.AuthService) *AuthController {
	return &AuthController{service}
}

// Register Регистрация пользователя
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var userRequest request.UserRequest

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		log.Errorln("Error decoding request.")
		handler.WrapBadRequest(w, errors.New("Error decoding request."))
		return
	}

	cookie, err := c.service.Registration(userRequest)
	if err != nil {
		log.Errorln("Error registering user")
		handler.WrapBadRequest(w, errors.New("Error registering user"))
		return
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusCreated)
}

// Login Вход пользователя в систему
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var userRequest request.UserRequest

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		handler.WrapBadRequest(w, errors.New("Error decoding request"))
		return
	}

	cookie, err := c.service.Login(userRequest)
	if err != nil {
		log.Errorln("Error login user")
		handler.WrapBadRequest(w, errors.New("Error login user"))
		return
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusCreated)
}
