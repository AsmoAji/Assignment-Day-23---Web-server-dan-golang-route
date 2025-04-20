package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

const maxUploadSize = 2 << 20

func UploadImage(c *gin.Context) {
	id := c.Param("id")

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	if file.Size > maxUploadSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File terlalu besar. Maksimum 2MB"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak diperbolehkan"})
		return
	}

	filename := fmt.Sprintf("uploads/product_%s%s", id, ext)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File berhasil diunggah", "path": filename})
}

func DownloadImage(c *gin.Context) {
	id := c.Param("id")
	for ext := range allowedExtensions {
		path := fmt.Sprintf("uploads/product_%s%s", id, ext)
		if _, err := os.Stat(path); err == nil {
			c.File(path)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Gambar tidak ditemukan"})
}
