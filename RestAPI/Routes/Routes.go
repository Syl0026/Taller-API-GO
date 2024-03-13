package Routes

import (
	"tallerAPIgo/RestAPI/Controllers"

	"github.com/gin-gonic/gin"
)

// % Definiendo las rutas de nuestro servicio
func RegisterRoutes(route *gin.Engine) {
	alumnRoute := route.Group("/alumn")
	alumnRoute.POST("/find", Controllers.FindAlumn)
	alumnRoute.POST("/create", Controllers.CreateAlumn)
	alumnRoute.PUT("/update", Controllers.UpdateAlumn)    //! PUT para el update
	alumnRoute.DELETE("/delete", Controllers.DeleteAlumn) //$ DELETE para el delete
}
