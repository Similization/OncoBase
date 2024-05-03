package model

type Doctor struct {
	Id            int    `json:"id" db:"id"`
	FirstName     string `json:"first-name" db:"first_name"`
	MiddleName    string `json:"middle-name" db:"middle_name"`
	LastName      string `json:"last-name" db:"last_name"`
	Qualification string `json:"qualification" db:"qualification"`
	Phone         string `json:"phone" db:"phone"`
	UserId        int    `json:"user-id" db:"user_id"`
}
