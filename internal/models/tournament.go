package models

import (
	"github.com/google/uuid"
	"time"
)

type TournamentStatus string

const (
	TournamentStatusPending    TournamentStatus = "PENDING"
	TournamentStatusInProgress TournamentStatus = "IN_PROGRESS"
	TournamentStatusOnGoing    TournamentStatus = "ONGOING"
	TournamentStatusOpen       TournamentStatus = "OPEN"
	TournamentStatusClosed     TournamentStatus = "CLOSED"
	TournamentStatusCancelled  TournamentStatus = "CANCELLED"
	TournamentStatusFinished   TournamentStatus = "FINISHED"
)

type Tournament struct {
	ID         uuid.UUID        `json:"id"`
	Name       string           `json:"name"`
	Game       string           `json:"game"`
	Status     TournamentStatus `json:"status"`
	StartDate  time.Time        `json:"startDate"`
	EndDate    time.Time        `json:"endDate"`
	MaxPlayers int              `json:"maxPlayers"`
}
