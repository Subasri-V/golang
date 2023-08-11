package controllers
//logic will be written in controller
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c*gin.Context) {
	// name:=c.param("name")
	c.String(http.StatusOK,"Hello")
}


func Register(c*gin.Context){
	c.String(http.StatusOK,"Hi")
}

// func GetProfile(c*gin.Context){

// }