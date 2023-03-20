package security

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// BcryptEncryptor Шифровщик/Расшифровщик bcrypt
type BcryptEncryptor struct {
}

var log = logrus.New()

// NewBcryptEncryption Создание нового шифровщика/расшифровщика
func NewBcryptEncryption() *BcryptEncryptor {
	return &BcryptEncryptor{}
}

// EncryptUserPassword Шифрование нового пароля
func (*BcryptEncryptor) EncryptUserPassword(password string) ([]byte, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Errorln("Error encrypting password. Error: ", err)
		return nil, err
	}

	return passwordHash, nil
}

// ComparePasswordAndHash Сравнение пароля и хеша пароля
func (*BcryptEncryptor) ComparePasswordAndHash(passwordHash []byte, password string) error {
	return bcrypt.CompareHashAndPassword(passwordHash, []byte(password))
}
