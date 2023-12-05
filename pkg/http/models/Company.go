package models

type Company struct {
	CommonModel
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null;unique"`
	Description string `gorm:"not null"`
	Users       []User
}
