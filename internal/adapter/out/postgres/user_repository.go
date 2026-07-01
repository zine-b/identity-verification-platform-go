package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zineb-b/identity-verification-platform-go/internal/domain"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (
			id,
			email,
			password_hash,
			status,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.Status,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT 
			id,
			email,
			password_hash,
			status,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
	`

	var user domain.User

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
