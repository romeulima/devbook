package repository

import (
	"context"
	"fmt"

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

func (r *UserRepository) GetUsers(ctx context.Context, param string) ([]models.User, error) {
	param = fmt.Sprintf("%%%s%%", param)
	rows, err := r.Db.Query(ctx, `
		SELECT id, name, nick, email, created_at FROM users WHERE name ILIKE $1 OR nick ILIKE $1
	`, param)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		var user models.User

		if err = rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, userId int) (models.User, error) {
	var user models.User
	err := r.Db.QueryRow(ctx,
		`SELECT id, name, nick, email, created_at FROM users WHERE id = $1`,
		userId).Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id int, user *models.User) error {
	_, err := r.Db.Exec(ctx,
		`UPDATE users SET name = $1, nick = $2, email = $3 WHERE id = $4`,
		user.Name, user.Nick, user.Email, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	_, err := r.Db.Exec(ctx,
		`DELETE FROM users WHERE id = $1`,
		id)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := r.Db.QueryRow(ctx,
		`SELECT id, password FROM users WHERE email = $1`,
		email).Scan(&user.ID, &user.Password)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func NewRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}
