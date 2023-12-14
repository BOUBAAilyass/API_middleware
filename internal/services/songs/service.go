package songs

import (
	"Projet_Middleware/internal/models/songs"
	repository "Projet_Middleware/internal/repositories/songs"

	"github.com/sirupsen/logrus"
)

func InsertSongs(song models.Song) error {
	err := repository.CreateSong(song)
	if err != nil {
		logrus.Errorf("Erreur lors de l'insertion de la chanson : %s", err.Error())
		return err
	}
	return nil
}
