package comments

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"Projet_Middleware/internal/models/comment"

	"Projet_Middleware/internal/services/comment"
)

func InsertComment(w http.ResponseWriter, r *http.Request) {
	var newComment models.Comment
	err := json.NewDecoder(r.Body).Decode(&newComment)
	if err != nil {
		logrus.Errorf("Erreur lors de la lecture du corps de la requête : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		logrus.Errorf("Erreur lors de la génération de l'identifiant UUID : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = comments.CreateComment(newComment)
	if err != nil {
		logrus.Errorf("Erreur lors de la création du commentaire : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(newComment)
	_, _ = w.Write(response)
}
