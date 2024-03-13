package RestApi

import (
	"tallerAPIgo/RestAPI/Routes"

	"github.com/gin-gonic/gin"
)

func RegisterControllers() bool {
	route := gin.Default()
	Routes.RegisterRoutes(route)
	route.Run(":3000")
	//* Es el puerto
	return true
}
