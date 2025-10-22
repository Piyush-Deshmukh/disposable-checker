package server

import (
	"fmt"
	"net/http"

	"github.com/Piyush-Deshmukh/disposable-checker/internal/config"
	"github.com/Piyush-Deshmukh/disposable-checker/internal/handlers"
	"github.com/gin-gonic/gin"
)

func NewServer(cfg *config.Config) *http.Server {
	router := gin.Default() // includes logger + recovery middleware

	router.GET("/", func(c *gin.Context) {
		htmlString := `<h1>Welcome to Disposable Checker!</h1>`
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlString))
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.POST("/validate/email", handlers.ValidateEmailHandler)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}
}
