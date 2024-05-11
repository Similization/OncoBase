package model

import "github.com/guregu/null/v5"

type PatientDisease struct {
	Stage     null.String `json:"stage" db:"stage"`
	Diagnosis null.String `json:"diagnosis" db:"diagnosis" binding:"required"`
	Patient   null.Int    `json:"patient" db:"patient" binding:"required"`
	Disease   null.String `json:"disease" db:"disease" binding:"required"`
}
