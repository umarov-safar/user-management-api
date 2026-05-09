package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umarov-safar/user-management-api/internal/models"
	"github.com/umarov-safar/user-management-api/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request models.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	authResponse, err := h.authService.RegisterWithToken(c.Request.Context(), request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, authResponse)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request models.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	authResponse, err := h.authService.LoginWithToken(c.Request.Context(), request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authResponse)
}
func (h *AuthHandler) VerifyEmail(c *gin.Context) {
	token := c.Query("token") // достаём token из URL параметра

	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token required"})
		return
	}

	user, err := h.authService.VerifyEmail(c.Request.Context(), token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email verified successfully",
		"user":    user,
	})
}
