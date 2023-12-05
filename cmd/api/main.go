package main

import (
	"fmt"
	"gotest/pkg/db"
	"gotest/pkg/http/repository"
)

func main() {
	fmt.Println("Server is running")

	//calling the database method to run the program
	db.GETDB()

	//repository
	repository.Initialize()

	//creating user and role
	repository.CreateUserAndRole()

}
