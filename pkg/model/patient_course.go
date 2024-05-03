package model

type PatientCourse struct {
	Id        int    `json:"id" db:"phone"`
	Patient   int    `json:"patient" db:"patient"`
	Disease   string `json:"disease" db:"disease"`
	Course    string `json:"course" db:"course"`
	Doctor    int    `json:"doctor" db:"doctor"`
	BeginDate string `json:"begin-date" db:"begin_date"`
	EndDate   string `json:"end-date" db:"end_date"`
	Diagnosis string `json:"diagnosis" db:"diagnosis"`
}
