package comments

import (
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"Projet_Middleware/internal/services/comment"

	"github.com/go-chi/chi/v5"
)

func DeleteComment(w http.ResponseWriter, r *http.Request) {

	commentID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		logrus.Errorf("Erreur lors de la récupération de l'ID du commentaire : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = comments.DeleteComment(commentID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du commentaire : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
