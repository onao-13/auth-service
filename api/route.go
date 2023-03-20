package api

import (
	"github.com/gorilla/mux"
	"security-service/internals/app/controller"
)

// CreateRouter API сервиса
func CreateRouter(authController *controller.AuthController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/auth/registration", authController.Register).Methods("POST")
	r.HandleFunc("/api/v1/auth/login", authController.Login).Methods("POST")
	//r.HandleFunc("/api/v1/auth/reset-password", authController.ResetPassword).Methods("PORT")
	return r
}
