package route

import (
	"aplikasi_golang/repository"
	"fmt"

	"gorm.io/gorm"
)

type productRoute struct {
	db *gorm.DB
}

func NewProductRoute(db *gorm.DB) *productRoute {
	return &productRoute{db}
}

func (r *productRoute) DeclareProductRoute() {
	productRepository := repository.NewProductRepository(r.db)

	code, err := productRepository.GenerateCode()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Kode tergenerate => %s", code)
}
