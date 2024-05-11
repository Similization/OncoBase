package model

type PatientCourse struct {
	Id        int    `json:"id" db:"phone" binding:"required"`
	Patient   int    `json:"patient" db:"patient" binding:"required"`
	Disease   string `json:"disease" db:"disease" binding:"required"`
	Course    string `json:"course" db:"course" binding:"required"`
	Doctor    int    `json:"doctor" db:"doctor" binding:"required"`
	BeginDate string `json:"begin-date" db:"begin_date" binding:"required"`
	EndDate   string `json:"end-date" db:"end_date"`
	Diagnosis string `json:"diagnosis" db:"diagnosis"`
}
