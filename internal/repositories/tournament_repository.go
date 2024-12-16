package repositories

import (
	"context"
	"database/sql"
	"kraken/internal/models"
)

type TournamentRepository interface {
	Create(ctx context.Context, tournament *models.Tournament) (*models.Tournament, error)
}

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (pr *PostgresRepo) Create(ctx context.Context, tournament *models.Tournament) (*models.Tournament, error) {
	//TODO implement me
	panic("implement me")
}
