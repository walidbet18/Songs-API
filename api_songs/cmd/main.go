package main

import (
	"songs/internal/helpers"

	"github.com/sirupsen/logrus"
)

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
