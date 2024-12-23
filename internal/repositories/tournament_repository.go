package repositories

import (
	"context"
	"database/sql"
	"kraken/internal/models"
)

type TournamentRepository interface {
	Create(ctx context.Context, tournament *models.Tournament) error
}

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (pr *PostgresRepo) Create(ctx context.Context, tournament *models.Tournament) error {
	//TODO implement me
	sql := `INSERT INTO tournaments (id, name, game, status, start_date, end_date, max_players) 
            VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := pr.db.ExecContext(ctx, sql, tournament.ID, tournament.Name, tournament.Game, tournament.Status,
		tournament.StartDate, tournament.EndDate, tournament.MaxPlayers)
	if err != nil {
		return err
	}

	return nil
}
