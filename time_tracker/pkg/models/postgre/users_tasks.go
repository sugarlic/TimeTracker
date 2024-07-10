package postgre

import (
	"gorm.io/gorm"
	"test.com/pkg/models"
)

type UserTasksService interface {
	Delete(id int) error
	Create(people *models.People) error
}

type UserTasksModel struct {
	DB *gorm.DB
}

func (m *UserTasksModel) Create(people *models.People) error {
	userTask := &models.UserTask{
		Name:       people.Name,
		Surname:    people.Surname,
		Patronymic: people.Patronymic,
		Address:    people.Address,
		TaskId:     1,
	}

	result := m.DB.Create(userTask)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *UserTasksModel) Delete(id int) error {
	// var user models.User
	// result := m.DB.First(&user, "id = ?", id)
	// if result.Error != nil {
	// 	return result.Error
	// }

	// result = m.DB.Delete(&user)
	// if result.Error != nil {
	// 	return result.Error
	// }

	return nil
}
