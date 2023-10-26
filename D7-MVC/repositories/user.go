package repositories

import (
	"MVC/configs"
	"MVC/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := configs.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user *models.User) error {
	if err := configs.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(id int) (interface{}, error) {
	var user models.User
	if err := configs.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *models.User) error {
	if err := configs.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	if err := configs.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
