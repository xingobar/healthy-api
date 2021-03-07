package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

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

func (w *Weight) GetWeight(deviceId string, id string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("device_id = ? AND id = ?", deviceId, id)
	}
}