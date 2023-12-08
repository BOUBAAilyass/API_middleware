package main

import (
	"Projet_Middleware/internal/comments"
	"Projet_Middleware/internal/helpers/comment"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func main() {

	router := chi.NewRouter()
	// comments------------------------------------------------------------------------------------------------------------------
	router.Post("/comments", comments.InsertComment)

	logrus.Info("[INFO] Web server started. Now listening on *:8084")
	logrus.Fatalln(http.ListenAndServe(":8084", router))

}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	comment_schemes := []string{
		`CREATE TABLE IF NOT EXISTS comments (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            music_id INTEGER,
            user_id INTEGER,
            content TEXT,
			rating REAL CHECK (rating BETWEEN 0 AND 5)
			
            
        );`,
	}

	for _, comment_scheme := range comment_schemes {
		if _, err := db.Exec(comment_scheme); err != nil {
			logrus.Fatalln("Could not generate comments table ! Error was : " + err.Error())
		}
	}

	helpers.CloseDB(db)
}
