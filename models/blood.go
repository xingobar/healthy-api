package models

import "time"

type Blood struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement" db:"id"`
	DeviceId string `json:"device_id" db:"device_id"`
	Pulse int `json:"pulse" db:"pulse"`
	Diastolic float32 `json:"diastolic" db:"diastolic"`
	Systolic float32 `json:"systolic" db:"systolic"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (Blood) TableName() string {
	return "bloods"
}