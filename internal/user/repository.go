package user

import (
	"context"
	"database/sql"
)

type Repository interface {
	SaveUser(context.Context, *User) error
	FindByID(context.Context, int) (*User, error)
	FindAll(context.Context) ([]*User, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (p PostgresRepository) SaveUser(ctx context.Context, user *User) error {
	//TODO implement me
	panic("implement me")
}

func (p PostgresRepository) FindByID(ctx context.Context, i int) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostgresRepository) FindAll(ctx context.Context) ([]*User, error) {
	//TODO implement me
	panic("implement me")
}
