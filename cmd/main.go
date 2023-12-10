package main

import (
	"Projet_Middleware/internal/comments"
	"Projet_Middleware/internal/helpers"

	"Projet_Middleware/internal/users"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func main() {

	router := chi.NewRouter()
	// comments------------------------------------------------------------------------------------------------------------------
	router.Post("/comments", comments.InsertComment)
	router.Get("/comments", comments.GetComments)
	router.Get("/comments/{id}", comments.GetComment)
	router.Put("/comments/{id}", comments.UpdateComment)
	router.Delete("/comments/{id}", comments.DeleteComment)

	//users------------------------------------------------------------------------------------------------------------------
	router.Post("/users", users.CreateUser)
	router.Get("/users", users.GetAllUsers)
	router.Get("/users/{id}", users.GetUserByID)

	logrus.Info("[INFO] Web server started. Now listening on *:8084")
	logrus.Fatalln(http.ListenAndServe(":8084", router))

}

func init() {
	//users------------------------------------------------------------------------------------------------------------------
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening users database : %s", err.Error())
	}

	if _, e := db.Exec("PRAGMA foreign_keys = ON;"); e != nil {
		logrus.Fatalln("Could not enable foreign keys ! Error was : " + e.Error())
	}

	users_schemes := []string{
		`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT	NOT NULL,
		password TEXT	NOT NULL,
		email TEXT NOT NULL
		
	);`,
	}

	for _, users_scheme := range users_schemes {
		if _, err := db.Exec(users_scheme); err != nil {
			logrus.Fatalln("Could not generate users table ! Error was : " + err.Error())
		}
	}

	// comments------------------------------------------------------------------------------------------------------------------

	comment_schemes := []string{
		`CREATE TABLE IF NOT EXISTS comments (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            music_id INTEGER NOT NULL,
            user_id INTEGER NOT NULL,
            content TEXT NOT NULL,
			rating REAL CHECK (rating BETWEEN 0 AND 5),
			FOREIGN KEY (content) REFERENCES users(username)
			
            
        );`,
	}

	for _, comment_scheme := range comment_schemes {
		if _, err := db.Exec(comment_scheme); err != nil {
			logrus.Fatalln("Could not generate comments table ! Error was : " + err.Error())
		}
	}

	helpers.CloseDB(db)

}
