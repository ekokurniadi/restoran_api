package handler

import (
	"aplikasi_golang/formatter"
	"aplikasi_golang/input"
	"aplikasi_golang/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetAllDataUser(c *gin.Context) {
	users, err := h.userService.GetAllDataUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "data": users})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mendapatkan data User",
		"data":    formatter.FormatManyUser(users),
	})
}

func (h *userHandler) SaveDataUser(c *gin.Context) {
	var userInput input.UserInput
	err := c.ShouldBindJSON(&userInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "data": ""})
		return
	}

	insertedUser, err := h.userService.SaveDataUser(userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "data": insertedUser})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menambahkan data User",
		"data":    formatter.FormatSingleUser(insertedUser),
	})

}

func (h *userHandler) DeleteUser(c *gin.Context) {
	var userID input.UserInputID
	err := c.ShouldBindUri(&userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "data": ""})
		return
	}

	deletedUser, err := h.userService.DeleteUser(userID.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "data": deletedUser})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menghapus data User",
		"data":    deletedUser,
	})

}
func (h *userHandler) UpdateUser(c *gin.Context) {
	var userID input.UserInputID
	err := c.ShouldBindUri(&userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "data": ""})
		return
	}
	var userData input.UserInput
	err = c.ShouldBindJSON(&userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "data": ""})
		return
	}

	updatedUser, err := h.userService.UpdateUser(userID.Id, userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "data": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil Merubah data user",
		"data":    formatter.FormatSingleUser(updatedUser),
	})
}

func (h *userHandler) GetUserByID(c *gin.Context) {
	var userID input.UserInputID
	err := c.ShouldBindUri(&userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "data": ""})
		return
	}
	user, err := h.userService.GetUserByID(userID.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil Mendapatkan data user",
		"data":    formatter.FormatSingleUser(user),
	})
}
