package routes

import (
	"github.com/anushkapandey/image_uploader_backend/controllers"
	"github.com/anushkapandey/image_uploader_backend/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, userDB *model.UserFile) {
	router.Use(cors.Default())

	router.POST("/login", controllers.Login(userDB))
	router.POST("/signup", controllers.SignUp(userDB))
	router.POST("/upload-image", controllers.UploadImage())
	router.POST("/get-images", controllers.GetImages()) //to get all images
	// //edit image
	// router.POST("/edit-image", controllers.EditImage())
}
