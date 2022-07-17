package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type appRoute struct {
	db *gorm.DB
}

func NewAppRoute(db *gorm.DB) *appRoute {
	return &appRoute{db}
}

func (r *appRoute) DeclareAppRoute() {
	router := gin.Default()
	api := router.Group("/api/v1")

	NewUserRoute(r.db, api).DeclareUserRoute()
	NewProductRoute(r.db).DeclareProductRoute()
	router.Run()

}
