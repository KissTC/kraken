package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"kraken/internal/models"
	"kraken/internal/repositories"
	"kraken/internal/responses"
	"time"
)

type TournamentService struct {
	Tr repositories.TournamentRepository
}

func NewTournamentService(tournRepo repositories.TournamentRepository) *TournamentService {
	return &TournamentService{Tr: tournRepo}
}

func (ts *TournamentService) CreateTournament(name, game string, startDate, endDate time.Time, MaxPlayers int) (*responses.CreateTournamentResponse, error) {
	// TODO: Logica de creacion de torneos

	// Crear un torneo por modelo
	tournament := &models.Tournament{
		ID:         uuid.New(),
		Name:       name,
		Game:       game,
		Status:     models.TournamentStatusPending,
		StartDate:  startDate,
		EndDate:    endDate,
		MaxPlayers: MaxPlayers,
	}

	err := ts.Tr.Create(context.Background(), tournament)
	if err != nil {
		return nil, err
	}

	response := &responses.CreateTournamentResponse{
		Response: fmt.Sprintf("Tournament created: %s", tournament.ID),
	}

	return response, nil
}
