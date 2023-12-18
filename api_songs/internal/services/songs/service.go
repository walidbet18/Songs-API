package songs

import (
	"database/sql"
	"net/http"
	"songs/internal/models"
	repository "songs/internal/repositories/songs"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllSongs() ([]models.Song, error) {
	var err error
	// Appel du repository
	songs, err := repository.GetAllSongs()
	// Gestion des erreurs
	if err != nil {
		logrus.Errorf("erreur lors de la récupération des chansons : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Une erreur s'est produite",
			Code:    500,
		}
	}

	return songs, nil
}

func GetSongByID(id uuid.UUID) (*models.Song, error) {
	song, err := repository.GetSongByID(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "Chanson introuvable",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("Erreur lors de la récupération de la chanson : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Une erreur s'est produite",
			Code:    http.StatusInternalServerError,
		}
	}

	return song, err
}

func AddSong(song *models.Song) (*models.Song, error) {
	err := repository.AddSong(song)
	if err != nil {
		logrus.Errorf("erreur lors de l'ajout de la chanson : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Échec de l'ajout de la chanson",
			Code:    http.StatusInternalServerError,
		}
	}

	return song, nil
}

func UpdateSong(song *models.Song) (*models.Song, error) {
	err := repository.EditSong(song)
	if err != nil {
		logrus.Errorf("erreur lors de la mise à jour de la chanson : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Échec de la mise à jour de la chanson",
			Code:    http.StatusInternalServerError,
		}
	}

	return song, nil
}

func DeleteSong(id uuid.UUID) error {
	err := repository.DeleteSong(id)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression de la chanson : %s", err.Error())
		return &models.CustomError{
			Message: "Échec de la suppression de la chanson",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}
