package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/romeulima/devbook/internal/config"
)

func Connect() (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s",
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbName)

	dbpool, err := pgxpool.New(context.Background(), connString)

	if err != nil {
		return nil, err
	}

	if err = dbpool.Ping(context.Background()); err != nil {
		dbpool.Close()
		return nil, err
	}

	return dbpool, nil
}
