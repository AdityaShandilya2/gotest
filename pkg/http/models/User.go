package models

import "time"

type User struct {
	CommonModel
	ID           uint      `gorm:"primaryKey"`
	AuthID       string    `gorm:"not null;unique"`
	CompanyID    uint      `gorm:"not null"`
	Company      Company   `gorm:"foreignKey:CompanyID"`
	Name         string    `gorm:"not null"`
	EmailID      string    `gorm:"not null;unique"`
	LastLoggedIn time.Time `gorm:"default:null"`
	Role         Role      `gorm:"foreignKey:RoleID"`
	RoleID       uint
}
