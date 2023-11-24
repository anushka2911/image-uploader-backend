package main

import (
	"log"
	"os"

	"github.com/anushkapandey/image_uploader_backend/model"
	"github.com/anushkapandey/image_uploader_backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	if _, err := os.Stat("users.txt"); os.IsNotExist(err) {
		file, err := os.Create("users.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	router := gin.Default()

	userDB, err := model.LoadUsersFromFile("users.txt")
	if err != nil {
		log.Fatal(err)
	}

	routes.AuthRoutes(router, userDB)

	port := ":8080"
	log.Printf("Server running on port %s...\n", port)

	err = router.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
