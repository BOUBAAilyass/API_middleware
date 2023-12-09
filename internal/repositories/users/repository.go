package users

import (
	"Projet_Middleware/internal/helpers"
	"Projet_Middleware/internal/models/users"

	"github.com/sirupsen/logrus"
)

func CreateUser(user models.User) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données users : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO users ( username, password , email) VALUES (?, ?, ?)",
		user.UserName, user.Password, user.Email)
	if err != nil {
		logrus.Errorf("Erreur lors de l'insertion de l'utilisateur dans la base de données : %s", err.Error())
		return err
	}

	return nil
}
