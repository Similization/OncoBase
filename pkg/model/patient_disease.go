package model

type PatientDisease struct {
	Stage     string `json:"stage" db:"stage"`
	Diagnosis string `json:"diagnosis" db:"diagnosis"`
	Patient   int    `json:"patient" db:"patient"`
	Disease   string `json:"disease" db:"disease"`
}
