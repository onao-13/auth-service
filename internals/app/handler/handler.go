package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

var log = logrus.New()

// WrapOk Обертка со статусом 200
func WrapOk(w http.ResponseWriter, cookie *http.Cookie) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

// WrapBadRequest Обретка со статусом 400
func WrapBadRequest(w http.ResponseWriter, err error) {
	var msg = map[string]any{
		"status": http.StatusBadRequest,
		"error":  err,
	}

	res, marsErr := json.Marshal(msg)
	if marsErr != nil {
		log.Errorln("Error marshalling response. Error: ", marsErr)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.Write(res)

	w.WriteHeader(http.StatusBadRequest)
}
