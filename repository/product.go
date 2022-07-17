package repository

import (
	"aplikasi_golang/entity"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GenerateCode() (string, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) GenerateCode() (string, error) {
	var product entity.Product
	query := "SELECT product_code from products ORDER BY id DESC LIMIT 1"

	err := r.db.Raw(query).Scan(&product.ProductCode).Error

	if err != nil {
		return product.ProductCode, err
	}

	productCode := ""
	sequence := 1

	if product.ProductCode == "" {
		productCode = fmt.Sprintf("PRD-000%d", sequence)
	} else {
		codeFromDatabase := strings.Split(product.ProductCode, "-")[1]
		convertToInt, err := strconv.Atoi(codeFromDatabase)
		if err != nil {
			return product.ProductCode, err
		}

		productCode = fmt.Sprintf("PRD-%04d", convertToInt+1)
	}

	return productCode, nil

}
