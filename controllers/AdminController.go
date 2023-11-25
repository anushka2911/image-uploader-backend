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
	"github.com/olahol/go-imageupload"
)

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create the uploads directory if it doesn't exist
		if err := os.Mkdir("uploads", 0777); err != nil && !os.IsExist(err) {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Parse the image file from the form
		fileHeader, err := c.FormFile("image")
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Create an Image object using go-imageupload
		img, err := imageupload.Process(c.Request, "image")
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Generate a unique filename for the uploaded image
		fileName := service.GenerateFilename(fileHeader.Filename)

		// Save the original image
		err = img.Save(fmt.Sprintf("uploads/%s", fileName))
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Open the ImageDetails.txt file for appending details
		file, err := os.OpenFile("ImageDetails.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer file.Close()

		// Extract labels from the form
		labels := c.PostForm("labels")

		// Record image details in the ImageDetails.txt file
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

		// Read content from ImageDetails.txt file
		content, err := ioutil.ReadFile("ImageDetails.txt")
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Split content into lines and parse image details
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
