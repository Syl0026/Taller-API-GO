package main //! SIEMPRE SE DEBE DE ESCRIBIR EL PAQUETE O CARPETA

import (
	"fmt"
	"tallerAPIgo/Models"
	"tallerAPIgo/Services/Services"
)

//% Ejemplo de función
// func main() {
// 	fmt.Println("Hola mundo")
// }

func main() {
	model := Models.AlumnModel{
		Name:     "Juanito",
		LastName: "Perez",
		Curp:     "00000000000000000",
		Email:    "juanpg@gmail.com",
		Phone:    "9999999999",
	}
	Services.CreateAlumn(&model)

	result, msg := Services.CreateAlumn(&model)
	fmt.Println(result)
	fmt.Println(msg)
}
