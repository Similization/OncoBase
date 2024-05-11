package model

import "github.com/guregu/null/v5"

type Patient struct {
	Id         null.Int    `json:"id" db:"id" binding:"required"`
	FirstName  null.String `json:"first-name" db:"first_name"`
	MiddleName null.String `json:"middle-name" db:"middle_name"`
	LastName   null.String `json:"last-name" db:"last_name"`
	BirthDate  null.String `json:"birth-date" db:"birth_date"`
	Sex        null.String `json:"sex" db:"sex"`
	SNILS      null.String `json:"snils" db:"snils"`
	UserId     null.Int    `json:"user-id" db:"user_id"`
	Phone      null.String `json:"phone" db:"phone"`
}
