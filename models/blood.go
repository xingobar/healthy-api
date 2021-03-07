package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

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

// 取得血壓
func (b *Blood) GetBlood(deviceId string, id string) func (db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("device_id = ? AND id = ?", deviceId, id)
	}
}