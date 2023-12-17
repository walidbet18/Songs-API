package songs

import (
	"encoding/json"
	"net/http"
	"songs/internal/models"
	"songs/internal/services/songs"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// GetSongByID récupère une chanson par son ID.
func GetSongByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songID, _ := ctx.Value("songID").(uuid.UUID)

	song, err := songs.GetSongByID(songID)
	if err != nil {
		logrus.Errorf("erreur : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(song)
	_, _ = w.Write(body)
	return
}
