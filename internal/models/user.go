package models

import "time"

type User struct {
	ID            string     `db:"id" json:"id"`
	Email         string     `db:"email" json:"email"`
	PasswordHash  string     `db:"password_hash" json:"-"`
	FirstName     string     `db:"first_name" json:"first_name"`
	LastName      string     `db:"last_name" json:"last_name"`
	Bio           *string    `db:"bio" json:"bio"`
	AvatarURL     *string    `db:"avatar_url" json:"avatar_url"`
	Role          string     `db:"role" json:"role"`
	EmailVerified bool       `db:"email_verified" json:"email_verified"`
	VerifiedAt    *time.Time `db:"verified_at" json:"verified_at"`
	CreatedAt     time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at" json:"updated_at"`
}

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	User        *User  `json:"user"`
}
