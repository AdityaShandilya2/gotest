package models

type Role struct {
	CommonModel
	ID          uint         `gorm:"primaryKey"`
	Code        string       `gorm:"not null;unique"`
	Name        string       `gorm:"not null"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
	Users       []User
}
