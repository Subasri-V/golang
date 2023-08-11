package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/api/hello",handleHello)
	http.HandleFunc("/Products",handleProducts)
	http.HandleFunc("/contactUs",handlecontactUs)
	fmt.Println("Server running on port :3000")
	log.Fatal(http.ListenAndServe(":3000",nil))
	
}

func handleHello(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
	if req.URL.Path!="/api/hello"{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"<h1>Page not found</h1>")
	} else {
		fmt.Fprintf(w,"<h1>Welcome to HTTP</h1>")
	}
	//fmt.Fprintf(w,"<h1>Welcome to HTTP</h1")
}

func handleProducts(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
	fmt.Fprintf(w,"<h1>Welcome to Productsssss</h1")
}

func handlecontactUs(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
	fmt.Fprintf(w,"<h1>subasri.v@netxd.com</h1")
}