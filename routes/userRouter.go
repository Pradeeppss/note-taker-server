package routes

import (
	controller "my-notes-backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/login",controller.Login())
	incomingRoutes.POST("/SignUp",controller.SignUp())
}