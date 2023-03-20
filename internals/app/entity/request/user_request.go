package request

// UserRequest JSON запрос для регистрации
type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Number   string `json:"number"`
}
