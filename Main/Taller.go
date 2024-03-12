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
		IdAlumn:   1,
		Name:      "Juanito",
		LastName:  "Prez",
		Curp:      "00000000000000000",
		Email:     "npg@gmail.com",
		Phone_cel: "9999999999",
	}
	//Services.UpdateAlumn(&model)
	result, msg := Services.UpdateAlumn(&model)
	fmt.Println(result)
	fmt.Println(msg)
}
