package songs

import (
	"Projet_Middleware/internal/helpers"
	"Projet_Middleware/internal/models/songs"

	"github.com/sirupsen/logrus"
)

func CreateSong(song models.Song) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO songs ( id, title, artist, album, year, duration ) VALUES (?, ?, ?, ?, ?, ?)",
		song.ID, song.Title, song.Artist, song.Album, song.Year, song.Duration)
	if err != nil {
		logrus.Errorf("Erreur lors de l'insertion de lachanson dans la base de données : %s", err.Error())
		return err
	}

	return nil
}

/* id INTEGER PRIMARY KEY AUTOINCREMENT,
title TEXT NOT NULL,
artist TEXT NOT NULL,
album TEXT NOT NULL,
year INTEGER NOT NULL,
duration INTEGER NOT NULL */
