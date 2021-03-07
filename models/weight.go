package models

import "time"

type Weight struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id" db:"id"`
	DeviceId string `db:"device_id" json:"device_id"`
	Number float32 `db:"number" json:"number"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (Weight) TableName() string {
	return "weights"
}