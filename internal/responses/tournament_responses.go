package responses

import (
	"github.com/google/uuid"
	"kraken/internal/models"
	"time"
)

type CreateTournamentResponse struct {
	ID         uuid.UUID               `json:"id"`
	Name       string                  `json:"name"`
	Game       string                  `json:"game"`
	Status     models.TournamentStatus `json:"status"`
	StartDate  time.Time               `json:"start_date"`
	EndDate    time.Time               `json:"end_date"`
	MaxPlayers int                     `json:"max_players"`
	CreatedAt  time.Time               `json:"created_at"`
}
