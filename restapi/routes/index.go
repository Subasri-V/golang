package routes
//all the routes will be 
import (
	"restapi/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router*gin.Engine){
	router.GET("/api/login",controllers.Login)
	router.POST("/api/login",controllers.Register)
	// router.POST("/api/register")
	// router.GET("/api/profile:id")
}