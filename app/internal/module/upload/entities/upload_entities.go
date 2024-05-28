package entities

import (
	"fmt"
	"log/slog"
	"time"

	dbconnect "github.com/bl4nc/gopload/pkg/db-connect"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type FileInfo struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	OriginalName string
	Path         string
	IsActive     bool       `gorm:"default:true"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	DeletedAt    *time.Time `gorm:"default:null"`
}

var db *gorm.DB

func init() {
	db := dbconnect.Connect()
	var err error
	if err = db.AutoMigrate(&FileInfo{}).Error; err != nil {
		slog.Error("Error migrating database: %v", err)
	}
}

func SaveFileInfo(id uuid.UUID, originalName string, path string) (uuid.UUID, error) {
	db := dbconnect.Connect()
	defer db.Close()

	fileInfo := FileInfo{
		ID:           id,
		OriginalName: originalName,
		Path:         path,
	}

	if err := db.Create(&fileInfo).Error; err != nil {
		return uuid.Nil, err
	}
	return fileInfo.ID, nil
}

func GetFileInfoByID(idArquivo uuid.UUID) (FileInfo, error) {
	db := dbconnect.Connect()
	defer db.Close()
	var fileInfo FileInfo
	if err := db.Where("id = ? AND is_active = ?", idArquivo, true).First(&fileInfo).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return FileInfo{}, fmt.Errorf("file not found")
		}
		return FileInfo{}, err
	}
	return fileInfo, nil
}

func UpdateFileInfoStatus(idArquivo uuid.UUID) error {
	db := dbconnect.Connect()
	defer db.Close()
	now := time.Now()
	if err := db.Model(&FileInfo{}).Where("id = ?", idArquivo).Update(map[string]interface{}{
		"is_active":  false,
		"deleted_at": now,
	}).Error; err != nil {
		return err
	}
	return nil
}
