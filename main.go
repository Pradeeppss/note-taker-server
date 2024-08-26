package main

import (
	"os"

	"my-notes-backend/database"
	"my-notes-backend/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var subjectCollection *mongo.Collection = database.OpenCollection(database.Client, "Subjects")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	
	routes.UserRoutes(router)
	
	routes.SubjectRoutes(router)
	routes.NotesRouter(router)	

	router.Run(":" + port)
}
