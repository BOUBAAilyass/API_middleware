package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Projet_Middleware/internal/models/users"
	"Projet_Middleware/internal/services/users"

	"github.com/go-chi/chi/v5"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := users.UpdateUser(userID, &updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
