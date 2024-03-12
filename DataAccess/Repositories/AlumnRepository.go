package Repositories //! SIEMPRE SE DEBE DE ESCRIBIR EL PAQUETE O CARPETA

import (
	"tallerAPIgo/Models"

	"gorm.io/gorm"
) //! Se importa solo

// % Create, Save y Delete son funciones ya creadas por go para sql

// & El error indica que devuelve el mensaje o tipo de error
func CreateAlumn(model *Models.AlumnModel, db *gorm.DB) error {
	return db.Create(model).Error
} //! No existe Try-catch
func UpdateAlumn(model *Models.AlumnModel, db *gorm.DB) error {
	return db.Save(model).Error
}

func DeleteAlumn(model *Models.AlumnModel, db *gorm.DB) error {
	return db.Delete(model).Error
}
