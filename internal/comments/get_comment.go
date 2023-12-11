package comments

import (
	"encoding/json"

	"Projet_Middleware/internal/models/comment"
	"Projet_Middleware/internal/services/comment"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func GetComment(w http.ResponseWriter, r *http.Request) {

	commentId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	comment, err := comments.GetCommentById(commentId)
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
	body, _ := json.Marshal(comment)
	_, _ = w.Write(body)
	return
}
