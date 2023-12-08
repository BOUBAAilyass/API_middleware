package comments

import (
	"Projet_Middleware/internal/models/comment"
	repository "Projet_Middleware/internal/repositories/comment"

	"database/sql"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

func CreateComment(comment models.Comment) error {
	err := repository.CreateComment(comment)
	if err != nil {
		logrus.Errorf("Erreur lors de la création du commentaire : %s", err.Error())
		return err
	}
	return nil
}

func GetAllComments() ([]models.Comment, error) {
	var err error
	// calling repository
	comments, err := repository.GetAllComments()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return comments, nil
}

func GetCommentById(id int) (*models.Comment, error) {
	comment, err := repository.GetCommentById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "comment not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving comments : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return comment, err
}
