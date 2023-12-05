package repository

import (
	"gotest/pkg/db"
	"gotest/pkg/http/models"
	"log"

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

func CreateUserAndRole() {

	//Create a role
	role := models.Role{
		Name: "Active",
		Code: "15590",
		Permissions: []models.Permission{{
			Name: "ticket_read",
			Code: "15457",
		}},
	}

	result := DB.Debug().Create(&role)

	if result.Error != nil {
		// log.Fatal(result)
		log.Println(error.Error(result.Error))
	}

	//Creating dummy company
	company := models.Company{
		Name:        "Tata",
		Description: "Car company",
	}

	res := DB.Debug().Create(&company)

	if res.Error != nil {
		// log.Fatal(result)
		log.Println(error.Error(res.Error))
	}

	//Creating dummy User
	users := []models.User{
		{Name: "rahul", RoleID: role.ID, AuthID: "test15", EmailID: "rahul@gmail.com", Company: company},
		{Name: "shannul", RoleID: role.ID, AuthID: "test16", EmailID: "shannul@gmail.com", Company: company},
	}

	result2 := DB.Debug().Create(&users)

	if result2.Error != nil {
		// log.Fatal(result)
		log.Println(error.Error(result2.Error))
	}

	result3 := DB.Debug().Model(&role).Where("ID = ?", role.ID).Update("Name", "Non-Active")

	if result3.Error != nil {
		// log.Fatal(result)
		log.Println(error.Error(result3.Error))
	}

	// Print the updated role and users
	var updatedRole models.Role
	DB.Preload("Users").First(&updatedRole, role.ID)
	println(updatedRole.Name) // SuperAdmin
	for _, user := range updatedRole.Users {
		println(user.Name)
	}

}
