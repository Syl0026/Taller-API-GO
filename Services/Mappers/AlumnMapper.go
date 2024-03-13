package Mappers

import (
	"tallerAPIgo/Models"
	"tallerAPIgo/Services/DTO/Request"
	"tallerAPIgo/Services/DTO/Response"
)

// $ Para no hacer el seteo desde el servicio
// ! Se recibe la información

// * Los * son punteros a otros directorios
func FindAlumnRequestToModel(request *Request.FindAlumnRequest, model *Models.AlumnModel){
	model.LastName = request.LastName
}

func CreateAlumnRequestToModel(request *Request.CreateAlumnRequest, model *Models.AlumnModel){
	model.LastName = request.LastName
	model.Name = request.Name
	model.Curp = request.Curp
	model.Email = request.Email
	model.Phone_cel =request.Phone_cel
	model.Suspended = request.Suspended
}

func UpdateAlumnRequestToModel(request *Request.UpdateAlumnRequest, model *Models.AlumnModel){
	model.IdAlumn = request.IdAlumn
	model.LastName = request.LastName
	model.Name = request.Name
	model.Curp = request.Curp
	model.Email = request.Email
	model.Phone_cel =request.Phone_cel
	model.Suspended = request.Suspended
}

func DeleteAlumnRequestToModel(request *Request.DeleteAlumnRequest, model *Models.AlumnModel){
	model.IdAlumn= request.IdAlumn
}

func AlumnModelToResponse(model *Models.AlumnModel, response *Response.AlumnResponse){
	response.IdAlumn =model.IdAlumn
	response.Name=model.Name
	response.LastName=model.LastName
	response.Curp=model.Curp
	response.Email=model.Email
	response.Phone_cel=model.Phone_cel
	response.Suspended=model.Suspended
}