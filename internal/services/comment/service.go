package comments

import (
	"Projet_Middleware/internal/models/comment"
	repository "Projet_Middleware/internal/repositories/comment"

	"github.com/sirupsen/logrus"
	//"log"
)

func CreateComment(comment models.Comment) error {
	err := repository.CreateComment(comment)
	if err != nil {
		logrus.Errorf("Erreur lors de la création du commentaire : %s", err.Error())
		return err
	}
	return nil
}
