package Services //! SIEMPRE SE DEBE DE ESCRIBIR EL PAQUETE O CARPETA

import (
	"tallerAPIgo/DataAccess/Queries"
	"tallerAPIgo/DataAccess/Repositories"
	infracestructure "tallerAPIgo/Infracestructure"
	"tallerAPIgo/Models"
	"tallerAPIgo/Services/DTO/Request"
	"tallerAPIgo/Services/DTO/Response"
	"tallerAPIgo/Services/Mappers"
	"tallerAPIgo/Services/Validators"
)

//! Lo anterior se coloca solo

// $ func-función nombreFunción(llave *Clase.Estructura) ValorRetorno {}
// $ [] indica un array y este es AlumnModel
func FindAlumn(request *Request.FindAlumnRequest) []Response.AlumnResponse {
	// Conección
	conn := infracestructure.ConnectDB()
	//! Cierra la base de datos (SIEMPRE SE DEBE HACER)
	sql, _ := conn.DB()
	defer sql.Close()

	response := []Response.AlumnResponse{}
	model := Models.AlumnModel{}
	Mappers.FindAlumnRequestToModel(request, &model)

	//!  Se va a ejecutar cuando termine el procedimiento del bloque
	//$Agregar el servicio...
	result := Queries.RetrieveAlumn(&model, conn)

	for _, item := range result {
		alumnResponse := Response.AlumnResponse{}
		Mappers.AlumnModelToResponse(&item, &alumnResponse)
		response = append(response, alumnResponse)
	}

	return response
}

// ! en donde esta nil debería de ser la conexión,
// & Se coloca nil como nulo
//& Solo se pone entre parentesis las variables de entorno cuando son más de dos

func CreateAlumn(request *Request.CreateAlumnRequest) (bool, string) {
	// Conección
	conn := infracestructure.ConnectDB()
	//! Cierra la base de datos (SIEMPRE SE DEBE HACER)
	sql, _ := conn.DB()
	defer sql.Close()

	model := Models.AlumnModel{}
	Mappers.CreateAlumnRequestToModel(request, &model)

	//* Validación de datos
	if ok, msg := Validators.CreateAlumnValidator(&model); !ok {
		return ok, msg
	}

	if err := Repositories.CreateAlumn(&model, conn); err != nil {
		return false, err.Error()
	}
	return true, "Success"
}

func UpdateAlumn(request *Request.UpdateAlumnRequest) (bool, string) {
	// Conección
	conn := infracestructure.ConnectDB()
	//! Cierra la base de datos (SIEMPRE SE DEBE HACER)
	sql, _ := conn.DB()
	defer sql.Close()

	// !Creación de model y Mappers

	model := Models.AlumnModel{}
	Mappers.UpdateAlumnRequestToModel(request, &model)

	//* Validación de datos
	if ok, msg := Validators.UpdateAlumnValidator(&model, conn); !ok {
		return ok, msg
	}

	// & Errores if
	if err := Repositories.UpdateAlumn(&model, conn); err != nil {
		return false, err.Error()
	}

	return true, "Success"
}

func DeleteAlumn(request *Request.DeleteAlumnRequest) (bool, string) {
	// Conección
	conn := infracestructure.ConnectDB()
	//! Cierra la base de datos (SIEMPRE SE DEBE HACER)
	sql, _ := conn.DB()
	defer sql.Close()

	model := Models.AlumnModel{}
	Mappers.DeleteAlumnRequestToModel(request, &model)

	//* Validación
	if ok, msg := Validators.DeleteAlumnValidator(model.IdAlumn, conn); !ok {
		return ok, msg
	}

	//& Errores if
	if err := Repositories.DeleteAlumn(&model, conn); err != nil {
		return false, err.Error()
	}

	return true, "Success"
}
