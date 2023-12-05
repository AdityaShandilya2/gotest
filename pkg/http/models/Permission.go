package models

type Permission struct {
	CommonModel
	ID    uint   `gorm:"primaryKey"`
	Code  string `gorm:"not null;unique"`
	Name  string `gorm:"not null"`
	Roles []Role `gorm:"many2many:role_permissions"`
}
