package helpers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func StartGinServer(server *gin.Engine) error {
	srv := &http.Server{
		Addr:           os.Getenv("SERVER_PORT"),
		Handler:        server,
		IdleTimeout:    10 * time.Second,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := srv.ListenAndServe()

	return err
}
