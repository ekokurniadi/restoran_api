package main

import (
	"aplikasi_golang/config"
	"aplikasi_golang/entity"
	"aplikasi_golang/route"
	"fmt"
)

func main() {
	fmt.Println("Belajar Golang")
	database, _ := config.DatabaseConfig()
	// memigrasi struct ke table database
	database.AutoMigrate(&entity.Product{})
	route.NewAppRoute(database).DeclareAppRoute()
}
