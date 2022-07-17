package entity

import "time"

type Product struct {
	ID           int
	ProductCode  string
	ProductName  string
	CategoryID   int
	CategoryName string
	ProductImage string
	IsActive     int
	IsHaveStock  int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
