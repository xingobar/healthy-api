package models

import "time"

type Blood struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
	DeviceId string `json:"device_id"`
	Pulse int `json:"pulse"`
	Diastolic float32 `json:"diastolic"`
	Systolic float32 `json:"systolic"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Blood) TableName() string {
	return "bloods"
}