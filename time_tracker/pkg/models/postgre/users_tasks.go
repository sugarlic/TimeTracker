package postgre

import (
	"time"

	"gorm.io/gorm"
	"test.com/pkg/models"
)

type UserTasksService interface {
	Delete(id int) error
	Create(people *models.People) error
	Update(user *models.UserTask) error
	StartTask(userId, task_id int) error
	EndTask(userId int) error
	GetList(filter map[string]interface{}, page, pageSize int) ([]*models.UserTask, error)
	GetUserWorkload(userId int, startDate, endDate time.Time) ([]*models.UserTask, error)
}

type UserTasksModel struct {
	DB *gorm.DB
}

func (m *UserTasksModel) GetUserWorkload(userId int, startDate, endDate time.Time) ([]*models.TaskWorkload, error) {
	var workloads []*models.TaskWorkload

	result := m.DB.Table("user_tasks").
		Select("user_tasks.task_id, tasks.name AS task_name, SUM(user_tasks.total_minutes) AS total_minutes").
		Joins("JOIN tasks ON user_tasks.task_id = tasks.id").
		Where("user_tasks.id = ? AND user_tasks.start_time BETWEEN ? AND ?", userId, startDate, endDate).
		Group("user_tasks.task_id, tasks.name").
		Order("total_minutes DESC").
		Scan(&workloads)

	if result.Error != nil {
		return nil, result.Error
	}

	return workloads, nil

}

func (m *UserTasksModel) GetList(filter map[string]interface{}, page, pageSize int) ([]*models.UserTask, error) {
	var res []*models.UserTask

	query := m.DB.Where(filter)
	query = query.Limit(pageSize).Offset((page - 1) * pageSize)

	result := query.Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}

	return res, nil
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

func (m *UserTasksModel) Update(user *models.UserTask) error {
	result := m.DB.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *UserTasksModel) StartTask(userId, task_id int) error {
	var user models.UserTask
	result := m.DB.First(&user, "id = ?", userId)
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

func (m *UserTasksModel) EndTask(userId int) error {
	var user models.UserTask
	result := m.DB.First(&user, userId)
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
