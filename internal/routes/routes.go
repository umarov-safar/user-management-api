package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/umarov-safar/user-management-api/internal/config"
	"github.com/umarov-safar/user-management-api/internal/handlers"
	"github.com/umarov-safar/user-management-api/internal/repositories"
	"github.com/umarov-safar/user-management-api/internal/services"
	"github.com/umarov-safar/user-management-api/internal/utils"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB, cfg *config.Config) {
	userRepo := repositories.NewUserRepository(db)
	jwtToken := &utils.JWToken{}
	jwtToken.SetSecretKey(cfg.JWTSecret).SetExpiration(cfg.JWTExpiration)

	emailService := services.NewEmailService(services.EmailConfig{
		Host:     cfg.EmailHost,
		Port:     cfg.EmailPort,
		User:     cfg.EmailUser,
		Password: cfg.EmailPassword,
		From:     cfg.EmailFrom,
	})

	emailRepo := repositories.NewEmailTokenRepository(db)

	authService := services.NewAuthService(userRepo, jwtToken, emailService, emailRepo, cfg.FrontendUrl)

	authHandler := handlers.NewAuthHandler(authService)

	public := router.Group("/api/v1")
	auth := public.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
	
	webAuth := router.Group("/auth")
	webAuth.GET("/verify-email", authHandler.VerifyEmail)
}
