package model

import "github.com/guregu/null/v5"

type UnitMeasure struct {
	Id        null.String `json:"id" db:"id" binding:"required"`
	Shorthand null.String `json:"shorthand"  db:"shorthand" binding:"required"`
	FullText  null.String `json:"full-text"  db:"full_text"`
	Global    null.String `json:"global"  db:"global"`
}
