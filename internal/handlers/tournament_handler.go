package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kraken/internal/services"
)

type TournamentHandler struct {
	tournamentService *services.TournamentService
}

func NewTournamentHandler(tournamentService *services.TournamentService) *TournamentHandler {
	return &TournamentHandler{
		tournamentService: tournamentService,
	}
}

func (th *TournamentHandler) GetTournaments(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
