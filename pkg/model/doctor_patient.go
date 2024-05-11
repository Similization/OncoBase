package model

type DoctorPatient struct {
	Patient int `json:"patient" db:"patient" binding:"required"`
	Doctor  int `json:"doctor" db:"doctor" binding:"required"`
}
