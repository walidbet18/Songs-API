package songs

import (
	"songs/internal/helpers"
	"songs/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	rows, err := db.Query("SELECT * FROM songs")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	songs := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.ID, &data.Title, &data.Artist, &data.Type, &data.Duration, &data.ReleaseYear)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}

	return songs, nil
}

func GetSongByID(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	row := db.QueryRow("SELECT * FROM songs WHERE id=?", id.String())

	var song models.Song
	err = row.Scan(&song.ID, &song.Title, &song.Artist, &song.Type, &song.Duration, &song.ReleaseYear)
	if err != nil {
		return nil, err
	}

	return &song, nil
}

func AddSong(song *models.Song) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	id, err := uuid.NewV4() // Générer un nouvel UUID
	if err != nil {
		return err
	}

	// Convertir UUID en chaîne avant de l'insérer dans la base de données
	_, err = db.Exec("INSERT INTO songs (id, title, artist, type, duration, releaseyear) VALUES (?, ?, ?, ?, ?, ?)",
		id.String(), song.Title, song.Artist, song.Type, song.Duration, song.ReleaseYear)
	if err != nil {
		return err
	}

	return nil
}

func EditSong(song *models.Song) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE songs SET title = ?, artist = ?, type = ?, duration = ?, releaseyear = ? WHERE id = ?",
		song.Title, song.Artist, song.Type, song.Duration, song.ReleaseYear, song.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSong(songID uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM songs WHERE id = ?", songID)
	if err != nil {
		return err
	}

	return nil
}
