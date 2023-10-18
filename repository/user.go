package repository

import (
	connectdb "github.com/Asad2730/RDS-DB-Auth/connectDB"
	"github.com/Asad2730/RDS-DB-Auth/model"
)

func CreateUser(user *model.User) (*model.User, error) {

	if err := connectdb.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(email string, user *model.User) (string, error) {

	if err := connectdb.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", err
	}

	if err := connectdb.DB.Save(&user).Error; err != nil {
		return "", err
	}

	return "Updated!", nil
}

func GetUsers(user []model.User) ([]model.User, error) {

	if err := connectdb.DB.Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
