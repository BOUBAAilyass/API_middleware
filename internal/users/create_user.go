package users

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"Projet_Middleware/internal/models/users"

	"Projet_Middleware/internal/services/users"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		logrus.Errorf("Erreur lors de la lecture du corps de la requête : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = users.CreateUsers(newUser)
	if err != nil {
		logrus.Errorf("Erreur lors de la création de l'utilisateur : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(newUser)
	_, _ = w.Write(response)
}
