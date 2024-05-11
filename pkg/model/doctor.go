package model

type Doctor struct {
	Id            int    `json:"id" db:"id" binding:"required"`
	FirstName     string `json:"first-name" db:"first_name" binding:"required"`
	MiddleName    string `json:"middle-name" db:"middle_name" binding:"required"`
	LastName      string `json:"last-name" db:"last_name" binding:"required"`
	Qualification string `json:"qualification" db:"qualification"`
	Phone         string `json:"phone" db:"phone"`
	UserId        int    `json:"user-id" db:"user_id"`
}
