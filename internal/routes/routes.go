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

	authService := services.NewAuthService(userRepo, jwtToken)

	authHandler := handlers.NewAuthHandler(authService)

	public := router.Group("/api/v1")
	auth := public.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
}
