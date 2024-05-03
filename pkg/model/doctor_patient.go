package model

type DoctorPatient struct {
	Patient int `json:"patient" db:"patient"`
	Doctor  int `json:"doctor" db:"doctor"`
}
