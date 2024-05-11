package model

import "github.com/guregu/null/v5"

type BloodCount struct {
	Id               null.String `json:"id" db:"id" binding:"required"`
	Description      null.String `json:"description" db:"description"`
	MinNormalValue   null.Float  `json:"min-normal-value" db:"min_normal_value" binding:"required"`
	MaxNormalValue   null.Float  `json:"max-normal-value" db:"max_normal_value" binding:"required"`
	MinPossibleValue null.Float  `json:"min-possible-value" db:"min_possible_value" binding:"required"`
	MaxPossibleValue null.Float  `json:"max-possible-value" db:"max_possible_value" binding:"required"`
	MeasureCode      null.String `json:"measure-code" db:"measure_code" binding:"required"`
}
