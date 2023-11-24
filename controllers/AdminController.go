package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/anushkapandey/image_uploader_backend/model"
	service "github.com/anushkapandey/image_uploader_backend/service"
	"github.com/gin-gonic/gin"
)

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := os.Mkdir("uploads", 0777); err != nil && !os.IsExist(err) {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		fileHeader, err := c.FormFile("image")
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		fileName := service.GenerateFilename(fileHeader.Filename)
		err = c.SaveUploadedFile(fileHeader, fmt.Sprintf("uploads/%s", fileName))
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		file, err := os.OpenFile("ImageDetails.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer file.Close()

		labels := c.PostForm("labels")

		imageDetails := fmt.Sprintf("File: %s, Labels: %s, Timestamp: %s\n", fileName, labels, time.Now().Format(time.RFC3339))
		if _, err := file.WriteString(imageDetails); err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "File uploaded successfully",
			"file":    fileName,
		})
	}
}

func GetImages() gin.HandlerFunc {
	return func(c *gin.Context) {
		var images []model.ImageDetail

		content, err := ioutil.ReadFile("ImageDetails.txt")
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		lines := strings.Split(string(content), "\n")

		for _, line := range lines {
			if line != "" {
				imageDetail := parseImageDetail(line)
				images = append(images, imageDetail)
			}
		}

		c.JSON(200, gin.H{
			"images": images,
		})
	}
}

func parseImageDetail(line string) model.ImageDetail {
	parts := strings.Split(line, ", ")

	filePart := strings.Split(parts[0], ": ")[1]
	labelsPart := strings.Split(parts[1], ": ")[1]

	return model.ImageDetail{
		FileName: strings.TrimSpace(filePart),
		Labels:   strings.TrimSpace(labelsPart),
	}
}
