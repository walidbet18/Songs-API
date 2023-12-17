package songs

import (
	"encoding/json"
	"net/http"
	"songs/internal/models"
	"songs/internal/services/songs"

	"github.com/sirupsen/logrus"
)

// GetAllSongs récupère toutes les chansons.
func GetSongs(w http.ResponseWriter, _ *http.Request) {
	// Appel du service
	songs, err := songs.GetAllSongs()
	if err != nil {
		// Journalisation de l'erreur
		logrus.Errorf("erreur : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// Écriture du code HTTP dans l'en-tête
			w.WriteHeader(customError.Code)
			// Écriture du message d'erreur dans le corps de la réponse
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(songs)
	_, _ = w.Write(body)
	return
}
