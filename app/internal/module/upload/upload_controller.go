package uploadmodule

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/bl4nc/gopload/internal/module/upload/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context) {
	userPath := c.PostForm("path")
	normalizedPath, err := NormalizePath(userPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid path"})
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get multipart form"})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid files list"})
		return
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	var uploadErrors []string
	var uploadResults []map[string]string

	for _, file := range files {
		wg.Add(1)
		go func(file *multipart.FileHeader) {
			defer wg.Done()
			fileUUID := uuid.New()
			extension := filepath.Ext(file.Filename)
			newFileName := fileUUID.String() + extension
			savePath := filepath.Join("uploads/", normalizedPath, newFileName)
			id, err := entities.SaveFileInfo(fileUUID, file.Filename, savePath)
			if err != nil {
				mu.Lock()
				uploadResults = append(uploadResults, map[string]string{"error": "Failed to save file info: " + file.Filename})
				mu.Unlock()
				return
			}
			if err := c.SaveUploadedFile(file, savePath); err != nil {
				mu.Lock()
				uploadResults = append(uploadResults, map[string]string{"error": "Failed to save file: " + file.Filename})
				mu.Unlock()
				return
			}
			mu.Lock()
			uploadResults = append(uploadResults, map[string]string{"uuid": id.String(), "originalName": file.Filename})
			mu.Unlock()
		}(file)
	}
	wg.Wait()
	if len(uploadErrors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": uploadErrors})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":       "Files uploaded successfully",
		"uploadResults": uploadResults,
	})
}

func DownloadFile(c *gin.Context) {
	idArquivo := c.Param("idArquivo")
	id, err := uuid.Parse(idArquivo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id format"})
		return
	}
	fileInfo, err := entities.GetFileInfoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	c.FileAttachment(fileInfo.Path, fileInfo.OriginalName)
}
