package Validators //! SIEMPRE SE DEBE DE ESCRIBIR EL PAQUETE O CARPETA

import (
	"tallerAPIgo/Models" //!Se importa solo

	"gorm.io/gorm"
)

// Cada servicio va a tener su validador
// Hay una librería para validar los datos

// & func es una función
// $ func NombreFunción(variable *Tipo(Clase).Variable(estructura)) (tipoReturn){}
func CreateAlumnValidator(model *Models.AlumnModel) (bool, string) {
	//? ok funciona como bandera, el mensaje se guarde en msg y se evalúa el ok
	if ok, msg := requiredData(model); !ok {
		return ok, msg
	}
	return true, ""
}

// $---------------------------------------------------------
func UpdateAlumnValidator(model *Models.AlumnModel, db *gorm.DB) (bool, string) {
	if model.IdAlumn <= 0 {
		return false, "IdAlum es requerido"
	}

	//! Falta validar si existe el id en la base de datos
	var count int64
	//% Hace una consulta select *
	db.Model(&Models.AlumnModel{}).Where("id_alumn = ?", model.IdAlumn).Count(&count)

	if count <= 0 {
		return false, "El IdAlumn no existe"
	}

	//? ok funciona como bandera, el mensaje se guarde en msg y se evalúa el ok
	if ok, msg := requiredData(model); !ok {
		return ok, msg
	}

	return true, ""
}

// $---------------------------------------------------------
func DeleteAlumnValidator(idAlumn uint, db *gorm.DB) (bool, string) {
	if idAlumn <= 0 {
		return false, "IdAlum es requerido"
	}

	//! Falta validar si existe el id en la base de datos
	var count int64
	//% Hace una consulta select *
	//& db.Base(&Carpeta.Objeto{}).Where("consulta = ?" llave.Atributo).Count(&variableCount)
	db.Model(&Models.AlumnModel{}).Where("id_alumno = ?", idAlumn).Count(&count)

	if count <= 0 {
		return false, "El IdAlumn no existe"
	}

	return true, ""
}

//% Si se coloca el nombre de la función en minúscula solo se puede utilizar dentro de esa clase.
//% Si se coloca en mayúscula se vuelve pública

func requiredData(model *Models.AlumnModel) (bool, string) {
	//% Estamos validando datos requeridos
	if model.Name == "" {
		return false, "Error: Nombre no válido, el nombre es requerido"
	}
	if len(model.LastName) == 0 {
		return false, "Error, Apellido no válido, el apellido es requerido"
	}
	if len(model.Curp) == 0 {
		return false, "Error, Curp inválido, el curp es requerido"
	}
	if len(model.Email) == 0 { //& las variables son: tipo.nombreVariable, el tipo es la llave para llamar a esa clase
		//& Este tipo se declaró al principio de la función.
		return false, "Error, Email inválido, el email es requerido"
	}
	if len(model.Phone) == 0 { //& len indica la longitud de la variable
		return false, "Error, Teléfono inválido, el teléfono es requerido"
	}

	//% Verificacion de longitud de campos
	if len(model.Name) > 150 {
		return false, "Error: Logitud máxima es 150"
	}
	if len(model.LastName) > 150 {
		return false, "Error: Logitud máxima es 150"
	}
	if len(model.Curp) > 18 {
		return false, "Error: Logitud máxima es 18"
	}
	if len(model.Email) > 200 {
		return false, "Error: Logitud máxima es 200"
	}
	if len(model.Phone) > 50 {
		return false, "Error: Logitud máxima es 50"
	}

	return true, ""
}
