package model

import "github.com/guregu/null/v5"

type PatientCourse struct {
	Id        null.Int    `json:"id" db:"phone" binding:"required"`
	Patient   null.Int    `json:"patient" db:"patient" binding:"required"`
	Disease   null.String `json:"disease" db:"disease" binding:"required"`
	Course    null.String `json:"course" db:"course" binding:"required"`
	Doctor    null.Int    `json:"doctor" db:"doctor" binding:"required"`
	BeginDate null.String `json:"begin-date" db:"begin_date" binding:"required"`
	EndDate   null.String `json:"end-date" db:"end_date"`
	Diagnosis null.String `json:"diagnosis" db:"diagnosis"`
}
