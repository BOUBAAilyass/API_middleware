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

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.UserID, &user.UserName, &user.Password, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	_ = rows.Close()

	return users, nil
}

func GetUserByID(id int) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM users WHERE id=?", id)
	helpers.CloseDB(db)

	var user models.User
	err = row.Scan(&user.UserID, &user.UserName, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id int, user *models.User) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE users SET username=?, password=?, email=? WHERE id=?",
		user.UserName, user.Password, user.Email, id)
	helpers.CloseDB(db)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM users WHERE id=?", id)
	helpers.CloseDB(db)
	if err != nil {
		return err
	}
	return nil
}
