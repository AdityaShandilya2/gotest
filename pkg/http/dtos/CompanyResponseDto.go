package dtos

import "gotest/pkg/http/models"

type CompanyResponseDto struct {
	Name        string
	Description string
	Users       []models.User
}
