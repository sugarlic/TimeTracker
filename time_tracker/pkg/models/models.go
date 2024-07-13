package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no current command")

type UserTask struct {
	ID           uint `gorm:"primaryKey"`
	Surname      string
	Name         string
	Patronymic   string
	Address      string
	TaskId       int
	StartTime    time.Time
	EndTime      time.Time
	TotalMinutes int
}

func (UserTask) TableName() string {
	return "user_tasks"
}

type Task struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
}

func (Task) TableName() string {
	return "tasks"
}

type People struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}

type TaskWorkload struct {
	TaskID       int    `json:"task_id"`
	TaskName     string `json:"task_name"`
	TotalMinutes int    `json:"total_minutes"`
}
