package handlers

import (
	"net/http"

	"github.com/Piyush-Deshmukh/disposable-checker/pkg/validator"
	"github.com/gin-gonic/gin"
)

type PhoneRequest struct {
	Number string `json:"number" binding:"required"`
}

func ValidatePhoneHandler(c *gin.Context) {
	var req PhoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid_request",
			"msg":   "Number field is required",
		})
		return
	}

	isValid, reason := validator.ValidatePhone(req.Number)

	c.JSON(http.StatusOK, gin.H{
		"number": req.Number,
		"valid":  isValid,
		"reason": reason,
	})
}
