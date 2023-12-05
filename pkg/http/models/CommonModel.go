package models

import "time"

type CommonModel struct {
	CreatedAt  time.Time `gorm:"autoCreateTime:true;column:created_on"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime:true;column:last_modified_on"`
	CreatedBy  string    `gorm:"column:created_by"`
	ModifiedBy string    `gorm:"column:last_modified_by"`
}
