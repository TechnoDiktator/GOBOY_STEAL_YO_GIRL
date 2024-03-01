package models

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Menu struct {
	ID			primitive.ObjectID 				`bson:"_id"`
	Name		string							`json:"name" validate:"required"`
	Category	string							`json:"category" validate:"required"` 	
	Start_Date	*time.Time						`json:"start_date"`
	End_date	*time.Time						`json:"end_date"`
	Created_At	time.Time						`json:"created_at"`
	Updated_At	time.Time						`json:"update_at"`
	Menu_id		string							`json:"food_id"`

}