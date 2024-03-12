package infracestructure

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// & condición a la base de datos
func ConnectDB() *gorm.DB {
	host := "localhost"
	user := "postgres"
	password := "12345"
	databaseName := "tallerapigo"
	port := "5432"

	//!-----------------
	//$ Crea un string ordenado
	strConn := fmt.Sprintf("host =%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		databaseName,
		port)

	//$ Realiza la conexión de postgres
	db, err := gorm.Open(postgres.Open(strConn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
