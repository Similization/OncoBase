package model

type ProcedureBloodCount struct {
	Value       string `json:"value" db:"value"`
	MeasureCode string `json:"measure-code" db:"measure_code"`
	Procedure   int    `json:"procedure" db:"procedure" binding:"required"`
	BloodCount  string `json:"blood-count" db:"blood_count" binding:"required"`
}
