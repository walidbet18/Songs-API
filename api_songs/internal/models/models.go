package models

import (
	"github.com/gofrs/uuid"
)

type Song struct {
	ID          *uuid.UUID `json:"id"`
	Title       string     `json:"title"`
	Artist      string     `json:"artist"`
	Type        string     `json:"type"` // Type peut être un genre musical par exemple
	Duration    string     `json:"duration"`
	ReleaseYear int        `json:"releaseyear"`
}
