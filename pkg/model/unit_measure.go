package model

type UnitMeasure struct {
	Id        string `json:"id" db:"id"`
	Shorthand string `json:"shorthand"  db:"shorthand"`
	FullText  string `json:"full-text"  db:"full_text"`
	Global    string `json:"global"  db:"global"`
}
