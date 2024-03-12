package Services //! SIEMPRE SE DEBE DE ESCRIBIR EL PAQUETE O CARPETA

import (
	"tallerAPIgo/DataAccess/Queries"
	"tallerAPIgo/DataAccess/Repositories"
	"tallerAPIgo/Models"
)

//! Lo anterior se coloca solo

// $ func-función nombreFunción(llave *Clase.Estructura) ValorRetorno {}
// $ [] indica un array y este es AlumnModel
func FindAlumn(request *Models.AlumnModel) []Models.AlumnModel {
	//strConn := ""

	//db, _ := gorm.Open("postgres", strConn)
	//db.DB().Ping()
	//$Agregar el servicio...
	service := Queries.RetrieveAlumn(request, nil)

	return service
}

// ! en donde esta nil debería de ser la conexión,
// & Se coloca nil como nulo
//& Solo se pone entre parentesis las variables de entorno cuando son más de dos

func CreateAlumn(request *Models.AlumnModel) (bool, interface{}) {
	Repositories.CreateAlumn(request, nil)
	return true, request
}

func UpdateAlumn(request *Models.AlumnModel) (bool, interface{}) {
	Repositories.UpdateAlumn(request, nil)
	return true, request
}

func DeleteAlumn(request *Models.AlumnModel) (bool, interface{}) {
	Repositories.DeleteAlumn(request, nil)
	return true, request
}
