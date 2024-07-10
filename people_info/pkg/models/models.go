package models

import "time"

type User struct {
	ID             int `gorm:"primaryKey"`
	PassportSerie  int
	PassportNumber int
	Surname        string `gorm:"type:varchar(50)"`
	Name           string `gorm:"type:varchar(50)"`
	Patronymic     string `gorm:"type:varchar(50)"`
	Address        string `gorm:"type:varchar(100)"`
	Created        time.Time
	Updated        time.Time
}

func (User) TableName() string {
	return "users"
}

type People struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}
