package postgre

import (
	"context"

	"gorm.io/gorm"
	"test.com/pkg/models"
)

type UserService interface {
	Delete(ctx context.Context, id int) error
}

type UserModel struct {
	DB *gorm.DB
}

func (m *UserModel) Delete(ctx context.Context, id int) error {
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
