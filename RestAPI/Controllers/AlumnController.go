package Controllers

import (
	"net/http"
	"tallerAPIgo/Services/DTO/Request"
	"tallerAPIgo/Services/Services"

	"github.com/gin-gonic/gin"
)

func FindAlumn(context *gin.Context) {
	request := Request.FindAlumnRequest{}

	// $ La pedición se bindia al request, automáticamente lo va a convertir
	context.Bind(&request)
	response := Services.FindAlumn(&request)

	context.JSON(http.StatusOK, response)
}

func CreateAlumn(context *gin.Context) {
	request := Request.CreateAlumnRequest{}

	context.Bind(&request)
	response, msg := Services.CreateAlumn(&request)

	if response {
		context.JSON(http.StatusOK, msg)
	} else {
		context.JSON(http.StatusConflict, msg)
	}
}

func UpdateAlumn(context *gin.Context) {
	request := Request.UpdateAlumnRequest{}

	//$ Lo de la pedición se reseta en con el Bind
	context.Bind(&request)
	response, msg := Services.UpdateAlumn(&request)

	if response {
		context.JSON(http.StatusOK, msg)
	} else {
		context.JSON(http.StatusConflict, msg)
	}
}

func DeleteAlumn(context *gin.Context) {
	request := Request.DeleteAlumnRequest{}

	context.Bind(&request)
	response, msg := Services.DeleteAlumn(&request)

	if response {
		context.JSON(http.StatusOK, msg)
	} else {
		context.JSON(http.StatusConflict, msg)
	}
}
