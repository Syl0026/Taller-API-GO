package Queries //! SIEMPRE SE DEBE DE ESCRIBIR EL PAQUETE O CARPETA

import (
	"tallerAPIgo/Models"

	"gorm.io/gorm"
)

// ! Se colocan funciones
// % Forzosamente tienes que utilizar todas las funciones o variables para correr
// % Los argumentos son name *Clase. 	[]Models.AlumModel es un return array
func RetrieveAlumn(model *Models.AlumnModel, db *gorm.DB) []Models.AlumnModel {
	//$ Definimos un arreglo de modelos para guardar los resultados
	arrayModel := []Models.AlumnModel{}
	//$ Obtenemos los where de acuerdo a los campos
	db = getWhere(model, db) //$Se llama a la función
	//$ Realizamos una consulta
	db.Find(&arrayModel) //$ Sirve como un select * from, el resultado se guarda en el arrayModel
	//! Para direccionar la variable se manda con el & donde esta la ubicación de memoria
	return arrayModel
}

// % gorm.DB devolverá una conexión
func getWhere(model *Models.AlumnModel, db *gorm.DB) *gorm.DB {
	db = db.Where("last_name = ?", model.LastName)

	//$ Retornamos el db con los where que hayan cumplido la condición
	return db
}
