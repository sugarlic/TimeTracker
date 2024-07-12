package postgre

import (
	"time"

	"gorm.io/gorm"
	"test.com/pkg/models"
)

type UserTasksService interface {
	Delete(id int) error
	Create(people *models.People) error
	StartTask(user_id, task_id int) error
	EndTask(user_id int) error
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

func (m *UserTasksModel) StartTask(user_id, task_id int) error {
	var user models.UserTask
	result := m.DB.First(&user, "id = ?", user_id)
	if result.Error != nil {
		return result.Error
	}

	user.TaskId = task_id
	user.StartTime = time.Now().UTC()
	user.TotalMinutes = 0

	result = m.DB.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *UserTasksModel) EndTask(user_id int) error {
	var user models.UserTask
	result := m.DB.First(&user, user_id)
	if result.Error != nil {
		return result.Error
	}

	user.EndTime = time.Now()
	user.TotalMinutes = int(time.Since(user.StartTime.UTC()).Minutes())

	result = m.DB.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *UserTasksModel) Delete(id int) error {
	var user models.UserTask
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
