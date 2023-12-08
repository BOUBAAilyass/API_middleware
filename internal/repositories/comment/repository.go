package comments

import (
	"Projet_Middleware/internal/helpers/comment"
	"Projet_Middleware/internal/models/comment"
	"fmt"

	"github.com/sirupsen/logrus"
)

func CreateComment(comment models.Comment) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO comments ( music_id, user_id, content, rating ) VALUES (?, ?, ?, ?)",
		comment.MusicID, comment.UserID, comment.Content, comment.Rating)
	if err != nil {
		logrus.Errorf("Erreur lors de l'insertion du commentaire dans la base de données : %s", err.Error())
		return err
	}

	return nil
}
