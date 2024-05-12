package model

import "github.com/guregu/null/v5"

type User struct {
	Id       null.Int    `json:"id" db:"id"`
	Email    null.String `json:"email" db:"email" validate:"required"`
	Password null.String `json:"password" db:"password" validate:"required"`
	Role     null.String `json:"role" db:"role" validate:"required"`
}

type AuthUser struct {
	Email    null.String `json:"email" binding:"required"`
	Password null.String `json:"password" binding:"required"`
}

// type User struct {
// 	Id       int    `json:"id" db:"id"`
// 	Email    string `json:"email" db:"email" validate:"required"`
// 	Password string `json:"password" db:"password" validate:"required"`
// 	Role     string`json:"role" db:"role" validate:"required"`
// }

// type AuthUser struct {
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }
