package server

type User struct {
	Id         int    `json:"-"`
	FirstName  string `json:"first_name" `
	MiddleName string `json:"middle_name" `
	LastName   string `json:"last_name" `
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
