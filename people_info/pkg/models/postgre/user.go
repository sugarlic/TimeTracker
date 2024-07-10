package postgre

import (
	"fmt"

	"gorm.io/gorm"
	"people_info.com/pkg/models"
)

type UserService interface {
	Delete(id int) error
	Latest() ([]models.User, error)
	Get(passportSerie, passportNumber int) (*models.User, error)
	Insert(user models.User) error
}

type UserModel struct {
	DB *gorm.DB
}

func (m *UserModel) Latest() ([]models.User, error) {
	var users []models.User
	result := m.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (m *UserModel) Delete(id int) error {
	var user models.User
	result := m.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	result = m.DB.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *UserModel) Get(passportSerie, passportNumber int) (*models.User, error) {
	var user *models.User
	result := m.DB.Take(&user, "passport_serie = ? AND passport_number = ? ", passportSerie, passportNumber)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (m *UserModel) Insert(user models.User) error {
	result := m.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println(user)
	return nil
}
