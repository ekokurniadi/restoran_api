package repository

import (
	"aplikasi_golang/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllDataUser() ([]entity.User, error)
	SaveDataUser(inputData entity.User) (entity.User, error)
	UpdateUser(inputUser entity.User) (entity.User, error)
	GetUserByID(ID int) (entity.User, error)
	DeleteUser(ID int) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAllDataUser() ([]entity.User, error) {
	user := []entity.User{}
	err := r.db.Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) SaveDataUser(inputData entity.User) (entity.User, error) {
	err := r.db.Create(&inputData).Error
	if err != nil {
		return inputData, err
	}

	return inputData, nil
}

func (r *userRepository) GetUserByID(ID int) (entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(inputUser entity.User) (entity.User, error) {
	err := r.db.Save(&inputUser).Error
	if err != nil {
		return inputUser, err
	}
	return inputUser, nil
}

func (r *userRepository) DeleteUser(ID int) (bool, error) {
	var user entity.User
	err := r.db.Where("id= ?", ID).Delete(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
