package model

type Course struct {
	Id          string  `json:"id" db:"id"`
	Period      int     `json:"period" db:"period"`
	Frequency   float32 `json:"frequency" db:"frequency"`
	Dose        float32 `json:"dose" db:"dose"`
	Drug        string  `json:"drug" db:"drug"`
	MeasureCode string  `json:"measure-code" db:"measure_code"`
}
