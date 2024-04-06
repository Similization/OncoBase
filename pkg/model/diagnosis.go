package model

type Diagnosis struct {
	Id          string `json:"id" db:"id"`
	Description string `json:"description" db:"description"`
}
