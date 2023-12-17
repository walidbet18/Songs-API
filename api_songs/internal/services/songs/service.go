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
				Message: "Chanson non trouvée",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("erreur lors de la récupération de la chanson : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Une erreur s'est produite",
			Code:    500,
		}
	}

	return song, err
}
