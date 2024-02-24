package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter , r *http.Request){
	if err := r.ParseForm(); err !=nil{
		fmt.Fprintf(w , "ParseForm() err %v" , err)
		return 
	}
	fmt.Fprintf(w , "POST Request Successful\n")
	name:=r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w , "Name = %s\n" , name)
	fmt.Fprintf(w , "Address = %s\n" , address)

}

func helloHandler(w http.ResponseWriter , r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w , "404 not found!" , http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w , "methid is not supported" , http.StatusNotFound)
		return 
	}
	fmt.Fprint(w , "hello!")

}

func main(){
	
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/" , fileserver)

	http.HandleFunc("/form" , formHandler)
	http.HandleFunc("/hello" ,helloHandler)
	http.HandleFunc("/formOpen", func(w http.ResponseWriter, r *http.Request) {
		// Serve the form.html file (modify path if needed)
		http.ServeFile(w, r, "./static/form.html")
	})
	fmt.Printf("Starting server at port 9090 \n")
	if err:= http.ListenAndServe(":9090" , nil); err!=  nil{
		log.Fatal(err)
	}
}






