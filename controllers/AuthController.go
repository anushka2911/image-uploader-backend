package controllers

import (
	"log"

	"github.com/anushkapandey/image_uploader_backend/model"
	"github.com/anushkapandey/image_uploader_backend/service"
	"github.com/gin-gonic/gin"
)

func Login(userDB *model.UserFile) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.LoginRequest
		err := c.ShouldBindJSON(&user)
		if err != nil {
			log.Printf("Error while parsing JSON: %v\n", err)
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		userExists, err := service.UserExists(userDB.Users, user.Email)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		if !userExists {
			c.JSON(400, gin.H{
				"error": "Cannot find user with given email",
			})
			return
		}

		role, err := service.ValidateCredentials(userDB.Users, user.Email, user.Password)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Login successful",
			"role":    role,
		})
	}
}

func SignUp(userDB *model.UserFile) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser model.User
		err := c.ShouldBindJSON(&newUser)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request format",
			})
			return
		}

		userExists, err := service.UserExists(userDB.Users, newUser.Email)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Internal server error",
			})
			return
		}

		if userExists {
			c.JSON(400, gin.H{
				"error": "User already exists",
			})
			return
		}

		err = service.AddUser(userDB.Users, newUser.Email, newUser.Password, newUser.Role)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Internal server error",
			})
			return
		}

		err = model.SaveUsersToFile("users.txt", userDB)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Internal server error",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "User added successfully",
		})
	}
}
