package handlers

import (
	"net/http"

	"github.com/Piyush-Deshmukh/disposable-checker/pkg/validator"
	"github.com/gin-gonic/gin"
)

type EmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// ValidateEmailHandler handles /validate/email route
func ValidateEmailHandler(c *gin.Context) {
	var req EmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid_request",
			"msg":   "Email field is required and must be valid",
		})
		return
	}

	isValid, reason := validator.ValidateEmail(req.Email)

	c.JSON(http.StatusOK, gin.H{
		"email":  req.Email,
		"valid":  isValid,
		"reason": reason,
	})
}
