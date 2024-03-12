package Services //! SIEMPRE SE DEBE DE ESCRIBIR EL PAQUETE O CARPETA

import (
	"tallerAPIgo/DataAccess/Queries"
	"tallerAPIgo/DataAccess/Repositories"
	infracestructure "tallerAPIgo/Infracestructure"
	"tallerAPIgo/Models"
	"tallerAPIgo/Services/Validators"
)

//! Lo anterior se coloca solo

// $ func-función nombreFunción(llave *Clase.Estructura) ValorRetorno {}
// $ [] indica un array y este es AlumnModel
func FindAlumn(request *Models.AlumnModel) []Models.AlumnModel {
	// Conección
	conn := infracestructure.ConnectDB()
	//! Cierra la base de datos (SIEMPRE SE DEBE HACER)
	sql, _ := conn.DB()
	defer sql.Close()
	//!  Se va a ejecutar cuando termine el procedimiento del bloque
	//$Agregar el servicio...
	service := Queries.RetrieveAlumn(request, conn)
	return service
}

// ! en donde esta nil debería de ser la conexión,
// & Se coloca nil como nulo
//& Solo se pone entre parentesis las variables de entorno cuando son más de dos

func CreateAlumn(request *Models.AlumnModel) (bool, string) {
	// Conección
	conn := infracestructure.ConnectDB()
	//! Cierra la base de datos (SIEMPRE SE DEBE HACER)
	sql, _ := conn.DB()
	defer sql.Close()

	//* Validación de datos
	if ok, msg := Validators.CreateAlumnValidator(request); !ok {
		return ok, msg
	}

	Repositories.CreateAlumn(request, conn)
	return true, "Success"
}

func UpdateAlumn(request *Models.AlumnModel) (bool, string) {
	// Conección
	conn := infracestructure.ConnectDB()
	//! Cierra la base de datos (SIEMPRE SE DEBE HACER)
	sql, _ := conn.DB()
	defer sql.Close()

	//* Validación de datos
	if ok, msg := Validators.UpdateAlumnValidator(request, conn); !ok {
		return ok, msg
	}
	Repositories.UpdateAlumn(request, conn)
	return true, "Success"
}

func DeleteAlumn(request *Models.AlumnModel) (bool, string) {
	// Conección
	conn := infracestructure.ConnectDB()
	//! Cierra la base de datos (SIEMPRE SE DEBE HACER)
	sql, _ := conn.DB()
	defer sql.Close()
	//* Validación
	if ok, msg := Validators.DeleteAlumnValidator(request.IdAlumn, conn); !ok {
		return ok, msg
	}
	Repositories.DeleteAlumn(request, conn)
	return true, "Success"
}
