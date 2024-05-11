package model

import "github.com/guregu/null/v5"

type User struct {
	Id       null.Int    `json:"id" db:"id"`
	Email    null.String `json:"email" binding:"required" db:"email"`
	Password null.String `json:"password" binding:"required"`
	Role     null.String `json:"role" binding:"required"`
}

type AuthUser struct {
	Email    null.String `json:"email" binding:"required"`
	Password null.String `json:"password" binding:"required"`
}
