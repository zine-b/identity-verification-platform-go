package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// pgx est une librairie Go pour se connecter à PostgreSQL.
// pgxpool veut dire : pool de connexions PostgreSQL.
// une seul cnx pour plusieurs requete

func NewPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, databaseURL)
}
