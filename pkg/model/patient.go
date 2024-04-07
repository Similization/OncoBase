package model

type Patient struct {
	Id         string `json:"id" db:"id"`
	FirstName  string `json:"first-name" db:"first_name"`
	MiddleName string `json:"middle-name" db:"middle_name"`
	LastName   string `json:"last-name" db:"last_name"`
	BirthDate  string `json:"birth-date" db:"birth_date"`
	Sex        string `json:"sex" db:"sex"`
	SNILS      string `json:"snils" db:"snils"`
	Phone      string `json:"phone" db:"phone"`
}
