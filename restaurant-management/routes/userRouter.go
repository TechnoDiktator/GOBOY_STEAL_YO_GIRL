package routes

import (
	"encoding/json"
	controller "restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)


func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users" , controller.GetUsers())
	incomingRoutes.GET("/users/:user_id" , controller.GetUsers())
	incomingRoutes.POST("/users/signup" , controller.SignUp())
	incomingRoutes.POST("users/login" , controller.Login())
}