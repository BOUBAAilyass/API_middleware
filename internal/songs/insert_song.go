package songs

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"Projet_Middleware/internal/models/songs"

	"Projet_Middleware/internal/services/songs"
)

func InsertSong(w http.ResponseWriter, r *http.Request) {
	var newSong models.Song
	err := json.NewDecoder(r.Body).Decode(&newSong)
	if err != nil {
		logrus.Errorf("Erreur lors de la lecture du corps de la requête : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = songs.InsertSongs(newSong)
	if err != nil {
		logrus.Errorf("Erreur lors de l'insertion de la chonson : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(newSong)
	_, _ = w.Write(response)
}
