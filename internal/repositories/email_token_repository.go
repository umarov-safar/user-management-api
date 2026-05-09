package repositories

import (
	"context"
	"database/sql"
	"time"
)

type EmailTokenRepository struct {
	db *sql.DB
}

func NewEmailTokenRepository(db *sql.DB) *EmailTokenRepository {
	return &EmailTokenRepository{db}
}

func (r *EmailTokenRepository) CreateToken(ctx context.Context, userId, token string, expiresAt time.Time) error {
	query := `
		INSERT INTO email_verification_tokens(user_id, token, expires_at) VALUES($1, $2, $3)
	`

	_, err := r.db.ExecContext(ctx, query, userId, token, expiresAt)

	return err
}

func (r *EmailTokenRepository) GetByToken(ctx context.Context, token string) (*string, error) {
	query := `
		SELECT user_id
		FROM email_verification_tokens
		WHERE token = $1 AND expires_at > NOW()
	`

	var userId string
	err := r.db.QueryRowContext(ctx, query, token).Scan(userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &userId, nil
}

func (r *EmailTokenRepository) DeleteToken(ctx context.Context, token string) error {
	query := "DELETE FROM email_verification_tokens WHERE token = $1"

	_, err := r.db.ExecContext(ctx, query, token)

	return err
}
