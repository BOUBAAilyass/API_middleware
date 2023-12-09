package users

import (
	"Projet_Middleware/internal/models/users"
	"Projet_Middleware/internal/services/users"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := users.GetAllUsers()
	if err != nil {
		// logging error
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// writing http code in header
			w.WriteHeader(customError.Code)
			// writing error message in body
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(users)
	_, _ = w.Write(body)
	return

}
