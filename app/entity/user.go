package entity

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100)"`
	Email     string `gorm:"type:varchar(100);uniqueIndex"`
	Password  string `gorm:"type:varchar(100)"`
	Roles     []Role `gorm:"many2many:user_roles"`
	CreatedAt time.Time
	UpdateAt  time.Time
}

type Role struct {
	ID         uint         `gorm:"primaryKey"`
	Name       string       `gorm:"type:varchar(100)"`
	Permission []Permission `gorm:"many2many:role_permission"`
	CreatedAt  time.Time
	UpdateAt   time.Time
}

type Permission struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdateAt  time.Time
}
