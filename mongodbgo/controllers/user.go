package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter  , r *http.Request){
	id :=p.ByName("id")
	if !bson.IsObjectIdHex((id)){
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex((id))
	u := models.User{}
	if err := uc.Session.DB("mongo-golang").C("users").FindId(oid).One() != nil {
		w.Writeheader(404)
		return 

		uj , err := json.Marshal(u)
		if err!= nil{
			fmt.Println(err)
		}

	}

	w.Header().Set("Content-Type" ,  "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w , "%s\n")
}



CreateUser



DeleteUser







