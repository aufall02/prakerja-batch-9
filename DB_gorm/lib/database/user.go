package database

import (
	"dborm/config"
	"dborm/models"
	"errors"

	"fmt"

	"gorm.io/gorm"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	result := config.DB.Find(&users)
	err := result.Error
	if result.RowsAffected < 1 {
		err = fmt.Errorf("data doesn't exist")
		return nil, err
	}
	if result.Error != nil {
		return nil, err

	}

	return users, nil
}

func GetUser(nama string) (interface{}, error) {
	var user models.User
	if err := config.DB.Where("nama = ?", nama).First(&user).Error; err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		return nil, err
	}
	return user, nil

}

func CreateUser(user models.User) (interface{}, error) {

	if e := config.DB.Create(&user).Error; e != nil {
		return nil, e
	}
	return user, nil

}

func DeleteUser(nama string) (interface{}, error) {

	var user models.User
	var err error
	result := config.DB.Delete(&user, "nama = ?", nama)

	err = result.Error

	if err != nil {
		return nil, err
	}

	if result.RowsAffected < 1 {
		err = fmt.Errorf("row with id=%s cannot be deleted because it doesn't exist", nama)
		return nil, err
	}

	return user, nil

}

func UpdateUser(user models.User, nama string) (interface{}, error) {
	if err := config.DB.Where("nama = ?", nama).First(&user).Error; err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		return nil, err
	}

	if e := config.DB.Save(&models.User{Email: user.Email, Nama: user.Nama, Password: user.Password}).Error; e != nil {
		return nil, e
	}
	return user, nil

}
