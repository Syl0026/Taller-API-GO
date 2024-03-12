package Validators //! SIEMPRE SE DEBE DE ESCRIBIR EL PAQUETE O CARPETA

import "tallerAPIgo/Models" //!Se importa solo

// Cada servicio va a tener su validador
// Hay una librería para validar los datos

//& func es una función
//$ func NombreFunción(variable *Tipo(Clase).Variable(estructura)) (tipoReturn){}
func CreateAlumnValidator(model *Models.AlumnModel) (bool, string) {
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

	return true, "" //$ Valores de return
}