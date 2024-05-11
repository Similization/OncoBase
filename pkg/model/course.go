package model

type Course struct {
	Id          string  `json:"id" db:"id" binding:"required"`
	Period      int     `json:"period" db:"period" binding:"required"`
	Frequency   float32 `json:"frequency" db:"frequency" binding:"required"`
	Dose        float32 `json:"dose" db:"dose" binding:"required"`
	Drug        string  `json:"drug" db:"drug" binding:"required"`
	MeasureCode string  `json:"measure-code" db:"measure_code" binding:"required"`
}
