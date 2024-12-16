package main

import (
	"github.com/gofiber/fiber/v2"
	db2 "kraken/internal/db"
	"kraken/internal/handlers"
	"kraken/internal/repositories"
	"kraken/internal/services"
	"log"
)

func main() {
	db, err := db2.GetConnection("postgres", "example", "postgres", "localhost", 5432)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pr := repositories.NewPostgresRepo(db)
	ts := services.NewTournamentService(pr)
	tournamentHandlers := handlers.NewTournamentHandler(ts)

	app := fiber.New()

	app.Get("/", tournamentHandlers.GetTournaments)

	app.Listen(":3000")
}
