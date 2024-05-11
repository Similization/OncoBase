package model

import "github.com/guregu/null/v5"

type ProcedureBloodCount struct {
	Value       null.String `json:"value" db:"value"`
	MeasureCode null.String `json:"measure-code" db:"measure_code"`
	Procedure   null.Int    `json:"procedure" db:"procedure" binding:"required"`
	BloodCount  null.String `json:"blood-count" db:"blood_count" binding:"required"`
}
