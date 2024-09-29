package userdao

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id        string    `gorm:"primaryKey,column:id"`
	Name      string    `gorm:"column:name"`
	Role      string    `gorm:"column:role"`
	Address   string    `gorm:"column:address"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
