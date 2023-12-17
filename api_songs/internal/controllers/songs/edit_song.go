package songs

import (
	"encoding/json"
	"net/http"
	"songs/internal/models"
	"songs/internal/services/songs"

	"github.com/sirupsen/logrus"
)

// ModifierChanson modifie les détails d'une chanson existante.
func EditSong(w http.ResponseWriter, r *http.Request) {
	var updatedSong models.Song                         // Pas de pointeur ici, juste la structure Song
	err := json.NewDecoder(r.Body).Decode(&updatedSong) // Utilisation de "&" pour obtenir l'adresse de la structure
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Requête invalide"))
		return
	}

	// Appeler la fonction de mise à jour de la chanson dans le service approprié
	_, err = songs.UpdateSong(&updatedSong) // Passer la référence de la structure mise à jour
	if err != nil {
		logrus.Errorf("Erreur lors de la modification de la chanson : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Erreur interne du serveur"))
		return
	}

	// Répondre avec un statut 200 OK si tout s'est bien passé
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Chanson mise à jour avec succès"))
}
