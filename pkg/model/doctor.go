package model

import "github.com/guregu/null/v5"

type Doctor struct {
	Id            null.Int    `json:"id" db:"id" binding:"required"`
	FirstName     null.String `json:"first-name" db:"first_name" binding:"required"`
	MiddleName    null.String `json:"middle-name" db:"middle_name" binding:"required"`
	LastName      null.String `json:"last-name" db:"last_name" binding:"required"`
	Qualification null.String `json:"qualification" db:"qualification"`
	Phone         null.String `json:"phone" db:"phone"`
	UserId        null.Int    `json:"user-id" db:"user_id"`
}
