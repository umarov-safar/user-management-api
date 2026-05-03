package repositories

import (
	"context"
	"database/sql"

	"github.com/umarov-safar/user-management-api/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (id, email, password_hash, first_name, last_name, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := ur.db.ExecContext(ctx, query,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, email, password_hash, first_name, last_name, bio, avatar_url, 
			role, email_verified, verified_at, created_at
		FROM users
		WHERE email = $1
	`
	user := &models.User{}
	err := ur.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Bio,
		&user.AvatarURL,
		&user.Role,
		&user.EmailVerified,
		&user.VerifiedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetById(ctx context.Context, id string) (*models.User, error) {
	query := `
		SELECT id, email, password_hash, first_name, last_name, bio, avatar_url, 
			role, email_verified, verified_at, created_at
		FROM users
		WHERE id = $1
	`
	user := &models.User{}
	err := ur.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Bio,
		&user.AvatarURL,
		&user.Role,
		&user.EmailVerified,
		&user.VerifiedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) Update(ctx context.Context, user *models.User) error {
	query := `
		UPDATE useres
		SET email = $1, password_hash = $2, first_name = $3, last_name = $4, bio = $5, 
		avatr_url = $6 role = $7, email_verified = $8, verified_at = $9, updated_at $10
		WHERE id=$11
	`

	_, err := ur.db.ExecContext(ctx, query,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.Bio,
		user.AvatarURL,
		user.Role,
		user.EmailVerified,
		user.VerifiedAt,
		user.UpdatedAt,
		user.ID,
	)

	return err
}
