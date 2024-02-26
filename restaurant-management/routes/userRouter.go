package routes

import (
	"encoding/json"
	"restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)


func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users" , controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id" , controllers.GetUsers())
	incomingRoutes.POST("/users/signup" , controllers.SignUp())
	incomingRoutes.POST("users/login" , controllers.Login())
}