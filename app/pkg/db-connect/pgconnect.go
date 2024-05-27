package dbconnect

import (
	"log/slog"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Connect() *gorm.DB {
	slog.Info("STRING" + os.Getenv("DB_CONNECTION_STRING"))
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic("Error connecting to database: " + err.Error())
	}
	return db
}
