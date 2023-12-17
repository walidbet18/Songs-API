package main

import (
	"net/http"
	"songs/internal/controllers/songs"
	"songs/internal/helpers"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/songs", func(r chi.Router) {
		r.Get("/", songs.GetSongs)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(songs.Ctx)
			r.Get("/", songs.GetSong)

		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))

}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database: %s", err.Error())
	}
	defer helpers.CloseDB(db)

	schemes := []string{
		`CREATE TABLE IF NOT EXISTS songs (
            id UUID PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            artist VARCHAR(100) NOT NULL,
            type VARCHAR(50) NOT NULL,
            duration VARCHAR(20) NOT NULL,
            release_year INT NOT NULL
        )`,
	}

	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table! Error was: " + err.Error())
		}
	}
}
