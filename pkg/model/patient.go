package model

import "database/sql"

type Patient struct {
	Id         int           `json:"id" db:"id"`
	FirstName  string        `json:"first-name" db:"first_name"`
	MiddleName string        `json:"middle-name" db:"middle_name"`
	LastName   string        `json:"last-name" db:"last_name"`
	BirthDate  string        `json:"birth-date" db:"birth_date"`
	Sex        string        `json:"sex" db:"sex"`
	SNILS      string        `json:"snils" db:"snils"`
	UserId     sql.NullInt64 `json:"user-id" db:"user_id"`
	Phone      string        `json:"phone" db:"phone"`
}
