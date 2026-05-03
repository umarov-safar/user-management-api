package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/umarov-safar/user-management-api/internal/models"
	"github.com/umarov-safar/user-management-api/internal/repositories"
	"github.com/umarov-safar/user-management-api/internal/utils"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) Register(ctx context.Context, email, password string) (*models.User, error) {
	existingUser, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	user := &models.User{
		ID:           uuid.New().String(),
		Email:        email,
		PasswordHash: passwordHash,
		Role:         models.RoleUser,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := s.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*models.User, error) {
	user, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.VerifyPassword(user.PasswordHash, password) {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
