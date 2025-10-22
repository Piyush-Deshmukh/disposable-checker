package server

import (
	"fmt"
	"net/http"

	"github.com/Piyush-Deshmukh/disposable-checker/internal/config"
	"github.com/gin-gonic/gin"
)

func NewServer(cfg *config.Config) *http.Server {
	router := gin.Default() // includes logger + recovery middleware

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}
}
