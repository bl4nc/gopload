package main

import (
	"os"

	"github.com/bl4nc/gopload/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	r := router.RoutesSetup()
	r.Run(":" + port)
}
