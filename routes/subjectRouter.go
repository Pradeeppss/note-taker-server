package routes

import (
	controller "my-notes-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SubjectRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/subjects",controller.GetAllSubjects())
	incomingRoutes.POST("/subject",controller.AddSubject())
}