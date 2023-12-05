package dtos

import "time"

type UserUpdateRequest struct {
	AuthID       uint
	CompanyID    uint
	Name         string
	EmailID      string
	LastLoggedIn time.Time
	RoleID       uint
}
