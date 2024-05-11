package model

import "github.com/guregu/null/v5"

type Course struct {
	Id          null.String `json:"id" db:"id" binding:"required"`
	Period      null.Int    `json:"period" db:"period" binding:"required"`
	Frequency   null.Float  `json:"frequency" db:"frequency" binding:"required"`
	Dose        null.Float  `json:"dose" db:"dose" binding:"required"`
	Drug        null.String `json:"drug" db:"drug" binding:"required"`
	MeasureCode null.String `json:"measure-code" db:"measure_code" binding:"required"`
}
