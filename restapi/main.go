package main

import (
	"fmt"
	"log"
	"restapi/constants"
	"restapi/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	//router get put post 
	router:=gin.Default()
	routes.AppRoutes(router)
	fmt.Println("Server running on port",constants.PORT)
	log.Fatal(router.Run(constants.PORT))
}