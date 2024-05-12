package model

import "github.com/guregu/null/v5"

type CourseProcedure struct {
	Id            null.Int    `json:"id" db:"id" binding:"required"`
	PatientCourse null.Int    `json:"patient-course" db:"patient_course" binding:"required"`
	Doctor        null.Int    `json:"doctor" db:"doctor" binding:"required"`
	BeginDate     null.String `json:"begin-date" db:"begin_date" binding:"required"`
	Period        null.Int    `json:"period" db:"period"`
	Result        null.String `json:"result" db:"result"`
}
