package routers

import (
	"github.com/gin-gonic/gin"
	"tiki/controllers"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	router.POST("/add", controllers.HandleNewUserRequest)
	router.POST("/login", controllers.HandleLoginRequest)
	router.POST("/changepassword", controllers.HandleChangePasswordRequest)
}

/*Run : run service with a port*/
func Run(port string) {
	router.Run(":" + port)
}
