package model

type User struct {
	Email    string `json:"email" binding:"required" db:"email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type AuthUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
