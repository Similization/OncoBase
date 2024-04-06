package model

type Disease struct {
	Id          string `json:"id" db:"id"`
	Description string `json:"description" db:"description"`
}
