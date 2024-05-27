package router

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/bl4nc/gopload/internal/middleware"
	uploadmodule "github.com/bl4nc/gopload/internal/module/upload"
	"github.com/gin-gonic/gin"
)

func RoutesSetup() *gin.Engine {
	slog.Info("run")

	r := gin.Default()

	username := os.Getenv("BASIC_AUTH_USERNAME")
	password := os.Getenv("BASIC_AUTH_PASSWORD")
	authorized := r.Group("/api", middleware.BasicAuthMiddleware(username, password))
	{
		authorized.POST("/upload", uploadmodule.UploadFile)
		authorized.GET("/file/:idArquivo", uploadmodule.DownloadFile)
	}
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ON"})
	})

	return r
}
