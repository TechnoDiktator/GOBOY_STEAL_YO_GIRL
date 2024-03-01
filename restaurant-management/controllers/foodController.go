package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"restaurant-management/models"
	"restaurant-management/database"
	"log"
	"net/http"
	
)




func GetFoods () gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}



func GetFood () gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx , bson.M{"food_id":foodId}).Decode(&food)

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError , gin.H{"error":"error occured while fetching the food item"})
		}
		c.JSON(http.StatusOK , food)

	}
}


func CreateFood () gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}



func UpdateFood () gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func round (num float64) int {

}


func toFixed(num float64 , precision int) float64 {

}








