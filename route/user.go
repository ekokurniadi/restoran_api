package route

import (
	"aplikasi_golang/handler"
	"aplikasi_golang/repository"
	"aplikasi_golang/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userRoute struct {
	db  *gorm.DB
	api *gin.RouterGroup
}

func NewUserRoute(db *gorm.DB, api *gin.RouterGroup) *userRoute {
	return &userRoute{db, api}
}

func (r *userRoute) DeclareUserRoute() {
	userRepository := repository.NewUserRepository(r.db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r.api.GET("/users", userHandler.GetAllDataUser)
	r.api.POST("/users", userHandler.SaveDataUser)
	r.api.GET("/users/:id", userHandler.GetUserByID)
	r.api.DELETE("/users/:id", userHandler.DeleteUser)
	r.api.PUT("/users/:id", userHandler.UpdateUser)
}
