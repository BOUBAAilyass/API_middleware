package users

import (
	"Projet_Middleware/internal/models/users"
	repository "Projet_Middleware/internal/repositories/users"

	"github.com/sirupsen/logrus"
)

func CreateUsers(user models.User) error {
	err := repository.CreateUser(user)
	if err != nil {
		logrus.Errorf("Erreur lors de la création de l'utilisateur : %s", err.Error())
		return err
	}
	return nil
}
