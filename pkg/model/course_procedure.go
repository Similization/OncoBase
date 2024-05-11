package model

type CourseProcedure struct {
	Id            int    `json:"id" db:"id" binding:"required"`
	PatientCourse int    `json:"patient-course" db:"patient_course" binding:"required"`
	Doctor        int    `json:"doctor" db:"doctor" binding:"required"`
	BeginDate     string `json:"begin-date" db:"begin_date" binding:"required"`
	Period        int    `json:"period" db:"period"`
	Result        int    `json:"result" db:"result"`
}
