package Repositories //! SIEMPRE SE DEBE DE ESCRIBIR EL PAQUETE O CARPETA

import (
	"tallerAPIgo/Models"

	"gorm.io/gorm"
) //! Se importa solo

// % Create, Save y Delete son funciones ya creadas por go para sql
func CreateAlumn(model *Models.AlumnModel, db *gorm.DB) {
	db.Create(model)
}

//! No existe Try-catch

func UpdateAlumn(model *Models.AlumnModel, db *gorm.DB) {
	db.Save(model)
}

func DeleteAlumn(model *Models.AlumnModel, db *gorm.DB) {
	db.Delete(model)
}
