package songs

import (
	"encoding/json"
	"net/http"
	"songs/internal/models"
	"songs/internal/services/songs"

	"github.com/sirupsen/logrus"
)

// AddSong ajoute une nouvelle chanson.
func AddSong(w http.ResponseWriter, r *http.Request) {
	// Décodez le corps de la requête JSON dans une structure de données models.Song
	var newSong models.Song
	err := json.NewDecoder(r.Body).Decode(&newSong)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Requête invalide"))
		return
	}

	// Appeler le service pour ajouter la nouvelle chanson
	_, err = songs.AddSong(&newSong)
	if err != nil {
		logrus.Errorf("Erreur lors de l'ajout de la chanson : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Erreur interne du serveur"))
		return
	}

	// Répondre avec un statut 201 Created si tout s'est bien passé
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Chanson ajoutée avec succès"))
}
