package model

type Diagnosis struct {
	Id          string `json:"id" db:"id" binding:"required"`
	Description string `json:"description" db:"description" binding:"required"`
}
