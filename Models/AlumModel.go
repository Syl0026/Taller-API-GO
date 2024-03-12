package Models //! SIEMPRE SE DEBE DE ESCRIBIR EL PAQUETE O CARPETA

//!No se manejan clases sino como estructuras

//% gorm es para declarar el tipo de variables de la base de datos
//& not null - no nulo
//& varchar - string
//& int - entero

//$ Declaración de estructura para sql
//$type NombreEstructura struc {} --- Es un arreglo de objeto (básicamente)

type AlumnModel struct { //@una estructura como clase
	//% SIEMPRE SE UTILIZA gorm:
	//& type: es el tipo de variable (int, varchar, bool, ...)
	//! Al guardarse se le otorga formalidad al código

	//% un entero especial que incrementa
	IdAlumn   uint   `gorm:"primaryKey;autoIncrement:true"` //Variable IdAlum como tipo entero no nulo
	LastName  string `gorm:"type:varchar(150);not null"`
	Name      string `gorm:"type:varchar(150);not null"`
	Phone     string `gorm:"type:varchar(50);not null"`
	Email     string `gorm:"type:varchar(200);not null"`
	Curp      string `gorm:"type:varchar(18);not null"`
	Suspended *bool  `gorm:"type:boolean"` //?  apunta a una dirección de memoria. Acepta nulos
}

//& indica el nombre de la tabla
func (alumn AlumnModel) TableName() string {
	return "alumn"
}
