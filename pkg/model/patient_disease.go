package model

type PatientDisease struct {
	Stage     string `json:"stage" db:"stage"`
	Diagnosis string `json:"diagnosis" db:"diagnosis" binding:"required"`
	Patient   int    `json:"patient" db:"patient" binding:"required"`
	Disease   string `json:"disease" db:"disease" binding:"required"`
}
