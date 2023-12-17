package songs

import (
	"net/http"
	"songs/internal/services/songs"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// DeleteSong est un contrôleur pour supprimer une chanson.
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de la chanson à supprimer depuis les paramètres de la requête
	ctx := r.Context()
	songID, ok := ctx.Value("songID").(uuid.UUID)
	if !ok || songID == uuid.Nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("ID de chanson manquant ou invalide"))
		return
	}

	// Appeler la fonction DeleteSong du service pour supprimer la chanson
	err := songs.DeleteSong(songID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression de la chanson : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Erreur interne du serveur"))
		return
	}

	// Répondre avec un statut 200 OK si tout s'est bien passé
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Chanson supprimée avec succès"))
}
