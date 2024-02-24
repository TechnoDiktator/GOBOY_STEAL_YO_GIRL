package models

import (
	"go-bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

var db  *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:""json":"name"`
	Authos string `json:"author"`
	Publication string `json:"publication"`

}


func init(){
	config.Connect()
	db = config.GetDB()

}
	