
package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"

)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/user/:id" ,uc.GetUser )
	r.POST("/user" , uc.CreateUser )
	r.DELETE("/user/:id" , uc.DeleteUser )
	http.ListenAndServe("localhost:9000" , r)

}

func getSession() *mgo.Session{
	s , _err := mgo.Dial("mongodb://localhost:27017")
	if err != nil{
		panic(err)
	}
	return s

}