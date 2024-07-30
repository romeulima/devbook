package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/romeulima/devbook/internal/models"
)

type UserRepository struct {
	Db *pgxpool.Pool
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	err := r.Db.QueryRow(ctx,
		`INSERT INTO users (name, nick, email, password) 
		VALUES ($1, $2, $3, $4) RETURNING id, created_at`,
		user.Name, user.Nick, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func NewRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}
