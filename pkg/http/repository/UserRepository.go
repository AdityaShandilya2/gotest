package repository

import (
	"gotest/pkg/db"
	"gotest/pkg/http/models"

	"gorm.io/gorm"
)

var DB *gorm.DB

// func init() {
// 	DB = db.GETDB()
// 	DB.AutoMigrate(&models.User{}) //check the db
// }

func Initialize() {
	DB = db.GETDB()

	//Clear all the tables inside DB
	//DB.Migrator().DropTable(&models.User{}, &models.Role{}, &models.Permission{}, &models.Company{})

	DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Company{})

}
