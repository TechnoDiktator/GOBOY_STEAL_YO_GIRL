package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"go-bookstore/pkg/utils"
	"go-bookstore/pkg/models"
)


var NewBook models.Book

func GetBook(w http.ResponseWriter , r *http.Request){
	newBooks:=models.GetAllNBooks()
	res , _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type" , "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}



func GetBookById(w http.ResponseWriter , r *http.Request){
	vars:= mux.Vars(r)
	bookId := vars["bookId"]
	ID , err := strconv.ParseInt(bookId , 0 , 0)
	if(err !=nill){
		fmt.Println("error while parsing")
	}
	bookDetails , _ := models.GetBookById(ID)

	res,_:=json.Marshal(bookDetails)
	w.Header().Set("Content-Type" ,"pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func CreateBook(w http.ResponseWriter  , r *http.Request){
	CreateBook := &models.Book()
	utils.ParseBody(r , CreateBook)
	b:= CreateBook.CreateBook()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)


}


func DeleteBook(w http.ResponseWriter   , r *http.Request){
	var  := mux.Vars(r)
	bookId := vars["bookId"]
	ID , err := strconv.ParseInt(bookId , 0 , 0)
	if(err != nil){
		fmt.Println("error while parsing")

	}
	book:= models.DeleteBook(ID)
	res,_:= json.Marshal(book)
	w.Header().Set("Content-Type" , "pkglication/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)
}

func UpdateBook(w http.ResponseWriter , r *http.request){
	var updateBook = &models.Book()
	utils.ParseBody(r , updateBook)
	varS:=-mux.Vars(r)
	bookId : vars["booksId"]
	ID , err :+ strconv.ParseInt(booksId , 0 , 0)

	bookDetails , db:=models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Author = updateBook.Name 
	} 
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}

	if(updateBook.Publicaton!=""){
		bookDetails.Publication = updateBook.Publications

	}

	db.Save(&bookDetails)
	res , _ :=json.Marshal(bookDetails)
	w.Headder().Set("Content-Type" , "pkglication")

	w.WriteHeader(http.StatusOK)
	w.Write(res)


}



















































































