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

func GetAllUsers() ([]models.User, error) {
	users, err := repository.GetAllUsers()
	if err != nil {
		logrus.Errorf("Erreur lors de la récupération des utilisateurs : %s", err.Error())
		return nil, err
	}
	return users, nil

}
func GetUserByID(id int) (*models.User, error) {
	user, err := repository.GetUserByID(id)
	if err != nil {
		logrus.Errorf("Erreur lors de la récupération de l'utilisateur : %s", err.Error())
		return nil, err
	}
	return user, nil
}

func UpdateUser(id int, user *models.User) error {
	err := repository.UpdateUser(id, user)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour de l'utilisateur : %s", err.Error())
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	err := repository.DeleteUser(id)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression de l'utilisateur : %s", err.Error())
		return err
	}
	return nil
}
