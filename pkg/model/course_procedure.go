package model

type CourseProcedure struct {
	Id            int    `json:"id" db:"id"`
	PatientCourse int    `json:"patient-course" db:"patient_course"`
	Doctor        int    `json:"doctor" db:"doctor"`
	BeginDate     string `json:"begin-date" db:"begin_date"`
	Period        int    `json:"period" db:"period"`
	Result        int    `json:"result" db:"result"`
}
