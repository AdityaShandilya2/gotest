package dtos

import "time"

type UserCreateRequest struct {
	AuthID       uint
	CompanyID    uint
	Name         string
	EmailID      string
	LastLoggedIn time.Time
	RoleID       uint
}
