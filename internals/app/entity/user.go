package entity

import "github.com/google/uuid"

// User Таблица с пользователями
type User struct {
	id       int64
	Uuid     uuid.UUID
	Name     string
	Password []byte
	Number   string
}
