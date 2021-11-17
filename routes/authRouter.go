package routes

import (
	controller "github.com/KeshikaGupta20/GoAuthJWT/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("user/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())

}
