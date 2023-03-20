package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"security-service/internals/app/entity"
	"security-service/internals/app/entity/request"
	"security-service/internals/app/repository"
	"security-service/internals/security"
	"security-service/internals/security/jwt"
)

// AuthService Сервисный слой для работы с пользователем entity.User
type AuthService struct {
	repository *repository.UserRepository
	encryptor  *security.BcryptEncryptor
	provider   *jwt.Provider
}

var log = logrus.New()

// NewAuthService Создание нового сервиса
func NewAuthService(
	repository *repository.UserRepository,
	encryptor *security.BcryptEncryptor,
	provider *jwt.Provider) *AuthService {
	return &AuthService{
		repository,
		encryptor,
		provider,
	}
}

// Registration Регистрация нового пользователя
func (s *AuthService) Registration(r request.UserRequest) (*http.Cookie, error) {
	registerErr := errors.New("Error registering u. Please try again later.")

	u, err := s.convertUserRequestToUser(r)
	if err != nil {
		return &http.Cookie{}, registerErr
	}

	err = s.repository.Save(u)
	if err != nil {
		return &http.Cookie{}, registerErr
	}

	return s.provider.GenerateUserToken(u.Uuid)
}

// Login Вход в систему
func (s *AuthService) Login(r request.UserRequest) (*http.Cookie, error) {
	u, err := s.repository.FindUserByNumber(r.Number)
	if err != nil {
		return &http.Cookie{}, errors.New("Error find u. Check your number")
	}

	loginErr := s.encryptor.ComparePasswordAndHash(u.Password, r.Password)
	if loginErr != nil {
		return &http.Cookie{}, errors.New("Password error. Check your password")
	}

	return s.provider.GenerateUserToken(u.Uuid)
}

// Конвертирование request.UserRequest в entity.User
func (s *AuthService) convertUserRequestToUser(r request.UserRequest) (entity.User, error) {
	var u = entity.User{}
	var err error

	u.Name = r.Name
	u.Number = r.Number
	u.Uuid = uuid.New()
	u.Password, err = s.encryptor.EncryptUserPassword(r.Password)
	if err != nil {
		log.Errorln("Error converting u password to byte. Error: ", err)
		return entity.User{}, err
	}

	return u, nil
}
