package users

import (
	"Projet_Middleware/internal/models/users"
	"Projet_Middleware/internal/services/users"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi/v5"
)

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := users.GetUserByID(userID)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(user)
	_, _ = w.Write(body)
	return
}
