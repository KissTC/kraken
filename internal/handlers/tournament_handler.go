package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kraken/internal/services"
	"time"
)

type TournamentHandler struct {
	tournamentService *services.TournamentService
}

func NewTournamentHandler(tournamentService *services.TournamentService) *TournamentHandler {
	return &TournamentHandler{
		tournamentService: tournamentService,
	}
}

type CreateTournamentRequest struct {
	Name       string    `json:"name"`
	Game       string    `json:"game"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	MaxPlayers int       `json:"max_players"`
}

func (th *TournamentHandler) GetAll(c *fiber.Ctx) error {
	//

	return c.SendString("Hello, World!")
}

func (th *TournamentHandler) Create(c *fiber.Ctx) error {
	// TODO: implement this method
	var request CreateTournamentRequest

	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid request")
	}

	response, err := th.tournamentService.CreateTournament(request.Name, request.Game, request.StartDate, request.EndDate, request.MaxPlayers)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
