package model

import "github.com/guregu/null/v5"

type DoctorPatient struct {
	Patient null.Int `json:"patient" db:"patient" binding:"required"`
	Doctor  null.Int `json:"doctor" db:"doctor" binding:"required"`
}
