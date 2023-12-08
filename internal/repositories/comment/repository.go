package comments

import (
	"Projet_Middleware/internal/helpers/comment"
	"Projet_Middleware/internal/models/comment"

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

func GetAllComments() ([]models.Comment, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM comments")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	comments := []models.Comment{}
	for rows.Next() {
		var data models.Comment
		err = rows.Scan(&data.ID, &data.MusicID, &data.UserID, &data.Content, &data.Rating)
		if err != nil {
			return nil, err
		}
		comments = append(comments, data)
	}

	_ = rows.Close()

	return comments, err
}

func GetCommentById(id int) (*models.Comment, error) {
	db, err := helpers.OpenDB()
	if err != nil {

		return nil, err
	}
	row := db.QueryRow("SELECT * FROM comments WHERE id=?", id)
	helpers.CloseDB(db)

	var comment models.Comment
	err = row.Scan(&comment.ID, &comment.MusicID, &comment.UserID, &comment.Content, &comment.Rating)

	if err != nil {

		return nil, err // Autres erreurs lors du scan
	}
	return &comment, err
}

func UpdateComment(comment *models.Comment) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE comments SET music_id=?, user_id=?, content=?, rating=? WHERE id=?",
		comment.MusicID, comment.UserID, comment.Content, comment.Rating, comment.ID)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour du commentaire dans la base de données : %s", err.Error())
		return err
	}

	return nil
}
