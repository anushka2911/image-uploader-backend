package routes

import (
	"github.com/anushkapandey/image_uploader_backend/controllers"
	"github.com/anushkapandey/image_uploader_backend/model"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, userDB *model.UserFile) {
	router.POST("/login", controllers.Login(userDB))
	router.POST("/signup", controllers.SignUp(userDB))
}
