package model

type UnitMeasure struct {
	Id        string `json:"id" db:"id" binding:"required"`
	Shorthand string `json:"shorthand"  db:"shorthand" binding:"required"`
	FullText  string `json:"full-text"  db:"full_text"`
	Global    string `json:"global"  db:"global"`
}
