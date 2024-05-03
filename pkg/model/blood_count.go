package model

type BloodCount struct {
	Id               string  `json:"id" db:"id"`
	Description      string  `json:"description" db:"description"`
	MinNormalValue   float32 `json:"min-normal-value" db:"min_normal_value"`
	MaxNormalValue   float32 `json:"max-normal-value" db:"max_normal_value"`
	MinPossibleValue float32 `json:"min-possible-value" db:"min_possible_value"`
	MaxPossibleValue float32 `json:"max-possible-value" db:"max_possible_value"`
	MeasureCode      string  `json:"measure-code" db:"measure_code"`
}
