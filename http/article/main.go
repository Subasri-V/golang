package main

import (
	"encoding/json"
	"io"
	"fmt"
	"log"
	"net/http"
)

//Go lang to json mapping
type Article struct{
	Id string `json:"Id"`
	Title string `json:"Title"`
	Content string `json:"Content"`
	Summary string `json:"Summary"`
	Author string `json:"Author"`

}

func main(){
	http.HandleFunc("/api/hello",handleHello)
	http.HandleFunc("/Products",handleProducts)
	http.HandleFunc("/contactUs",handlecontactUs)
	http.HandleFunc("/article",handleArticle)
	http.HandleFunc("/getArticles",getArticles)
	fmt.Println("Server running on port :3000")
	log.Fatal(http.ListenAndServe(":3000",nil))
	
}

func handleHello(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
	if req.URL.Path!="/api/hello"{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w,"<h1>Page not found</h1>")
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

func handleArticle(w http.ResponseWriter,req *http.Request){
	if req.Method=="POST"{
		reqBody,_:=io.ReadAll(req.Body)
		var post Article
		err:=json.Unmarshal(reqBody,&post)//json to go format

		post.Author="John"
		if err !=nil{
			fmt.Println("I go an Error")
			fmt.Fprintf(w,err.Error())
		} else {
			json.NewEncoder(w).Encode(post)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"Unable to process your request")
	}
}


func getArticles(w http.ResponseWriter,req *http.Request){
	
	if req.Method=="POST"{
		reqBody,_:=io.ReadAll(req.Body)
		var articles []Article
		err:=json.Unmarshal(reqBody,&articles)
		if err!=nil {
			fmt.Println("I go an Error")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w,err.Error())
		} else {
			fmt.Println(articles)
			fmt.Fprintf(w,"<h1>Success</h1>")
			articles=append(articles,Article{Id:"3",Title:"MIB"})
			json.NewEncoder(w).Encode(articles)
		}
	}
}