package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"security-service/internals/conf"
	"time"
)

// Provider Провайдер для работы с JWT
type Provider struct {
	secret string
}

var log = logrus.New()

// NewProvider Создание нового провайдера
func NewProvider(cfg conf.Config) *Provider {
	return &Provider{secret: cfg.JwtSecret}
}

// GenerateAdminToken Генерирование токена для администратора
func (p *Provider) GenerateAdminToken(userUuid uuid.UUID) (*http.Cookie, error) {
	token, err := p.generateToken(userUuid, "admin")
	cookie := saveJwtToCookie(token)
	return cookie, err
}

// GenerateUserToken Генерирование токена для пользователя
func (p *Provider) GenerateUserToken(userUuid uuid.UUID) (*http.Cookie, error) {
	token, err := p.generateToken(userUuid, "user")
	cookie := saveJwtToCookie(token)
	return cookie, err
}

// TODO: Сменить симметричный HS256 ключ на асиммитричный RS256 ключи
// Создание токена
func (p *Provider) generateToken(userUuid uuid.UUID, accessType string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userUuid,
			"access":  accessType,
			"exp":     time.Now().Add(1 * time.Hour).Unix(),
		})

	tokenString, err := token.SignedString([]byte(p.secret))
	if err != nil {
		log.Errorln("Error create token. Error: ", err)
		return "", err
	}

	return tokenString, nil
}

// Сохранение токена в куки
func saveJwtToCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:     "user_info",
		Value:    token,
		HttpOnly: true,
	}
}
