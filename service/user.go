package service

import (
	"aplikasi_golang/entity"
	"aplikasi_golang/input"
	"aplikasi_golang/repository"
	"errors"
)

type UserService interface {
	GetAllDataUser() ([]entity.User, error)
	SaveDataUser(inputData input.UserInput) (entity.User, error)
	GetUserByID(ID int) (entity.User, error)
	UpdateUser(ID int, inputData input.UserInput) (entity.User, error)
	DeleteUser(ID int) (bool, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) GetAllDataUser() ([]entity.User, error) {
	users, err := s.userRepository.GetAllDataUser()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *userService) SaveDataUser(inputData input.UserInput) (entity.User, error) {
	user := entity.User{}
	user.UserName = inputData.UserName
	user.Password = inputData.Password
	user.FirstName = inputData.FirstName
	user.LastName = inputData.LastName
	user.Gender = inputData.Gender

	var defaultAge int
	if inputData.Age == 0 {
		defaultAge = 17
	} else {
		defaultAge = inputData.Age
	}

	user.Age = defaultAge
	user.Address = inputData.Address

	var defaultRole string
	if inputData.Role == "" {
		defaultRole = "kasir"
	} else {
		defaultRole = inputData.Role
	}

	user.Role = defaultRole

	var defaultProfilePicture string
	if inputData.ProfilePicture == "" {
		defaultProfilePicture = "user_image.png"
	} else {
		defaultProfilePicture = inputData.ProfilePicture
	}
	user.ProfilePicture = defaultProfilePicture

	var defaultBlood string
	if inputData.Blood == "" {
		defaultBlood = "-"
	} else {
		defaultBlood = inputData.Blood
	}

	user.Blood = defaultBlood

	insertedUser, err := s.userRepository.SaveDataUser(user)
	if err != nil {
		return insertedUser, err
	}

	return insertedUser, nil
}

func (s *userService) GetUserByID(ID int) (entity.User, error) {
	user, err := s.userRepository.GetUserByID(ID)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, errors.New("user tidak ditemukan")
	}

	return user, nil
}

func (s *userService) UpdateUser(ID int, inputData input.UserInput) (entity.User, error) {
	user, err := s.userRepository.GetUserByID(ID)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, errors.New("user tidak ditemukan")
	}

	user.FirstName = inputData.FirstName
	user.LastName = inputData.LastName
	user.Age = inputData.Age
	user.Gender = inputData.Gender
	user.Password = inputData.Password
	user.Role = inputData.Role
	user.ProfilePicture = inputData.ProfilePicture
	user.UserName = inputData.UserName
	user.Address = inputData.Address
	user.Blood = inputData.Blood

	updatedUser, err := s.userRepository.UpdateUser(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}

func (s *userService) DeleteUser(ID int) (bool, error) {
	user, err := s.userRepository.GetUserByID(ID)
	if err != nil {
		return false, err
	}
	if user.Id == 0 {
		return false, errors.New("user tidak ditemukan")
	}

	deletedUser, err := s.userRepository.DeleteUser(ID)
	if err != nil {
		return false, err
	}

	return deletedUser, nil
}
