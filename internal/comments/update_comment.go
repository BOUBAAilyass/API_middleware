package comments

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"Projet_Middleware/internal/models/comment"
	"Projet_Middleware/internal/services/comment"

	"github.com/go-chi/chi/v5"
)

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	commentID, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		logrus.Errorf("Erreur lors de la récupération de l'ID du commentaire : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updatedComment models.Comment
	err = json.NewDecoder(r.Body).Decode(&updatedComment)
	if err != nil {
		logrus.Errorf("Erreur lors de la lecture du corps de la requête : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = comments.UpdateComment(commentID, updatedComment)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour du commentaire : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(updatedComment)
	_, _ = w.Write(response)
}
