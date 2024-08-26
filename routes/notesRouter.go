package routes

import (
	controller "my-notes-backend/controllers"

	"github.com/gin-gonic/gin"
)

func NotesRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/notes",controller.GetAllNotes())
	incomingRoutes.POST("/note",controller.AddNote())
	incomingRoutes.PUT("/note",controller.UpdateNote())

}