package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type Account interface {
}

type Authorization interface {
	CreateUser(user model.User) (string, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (*UserData, error)
}

type BloodCount interface {
	CreateBloodCount(bloodCount model.BloodCount) (model.BloodCount, error)
	GetBloodCountById(id string) (model.BloodCount, error)
	GetBloodCountList() ([]model.BloodCount, error)
	UpdateBloodCount(bloodCount model.BloodCount) (model.BloodCount, error)
	DeleteBloodCount(id string) error
}

type BloodCountValue interface {
	CreateBloodCountValue(bloodCountValue model.BloodCountValue) (model.BloodCountValue, error)
	GetBloodCountValueById(diseaseId, bloodCountId string) (model.BloodCountValue, error)
	GetBloodCountValueListByDisease(diseaseId string) ([]model.BloodCountValue, error)
	GetBloodCountValueListByBloodCount(bloodCountId string) ([]model.BloodCountValue, error)
	GetBloodCountValueList() ([]model.BloodCountValue, error)
	UpdateBloodCountValue(bloodCountValue model.BloodCountValue) (model.BloodCountValue, error)
	DeleteBloodCountValue(diseaseId, bloodCountId string) error
}

type Course interface {
	CreateCourse(course model.Course) (model.Course, error)
	GetCourseById(id string) (model.Course, error)
	GetCourseList() ([]model.Course, error)
	UpdateCourse(course model.Course) (model.Course, error)
	DeleteCourse(id string) error
}

type CourseProcedure interface {
	CreateCourseProcedure(courseProcedure model.CourseProcedure) (model.CourseProcedure, error)
	GetCourseProcedureById(id string) (model.CourseProcedure, error)
	GetCourseProcedureList() ([]model.CourseProcedure, error)
	UpdateCourseProcedure(courseProcedure model.CourseProcedure) (model.CourseProcedure, error)
	DeleteCourseProcedure(id string) error
}

type Diagnosis interface {
	CreateDiagnosis(diagnosis model.Diagnosis) (model.Diagnosis, error)
	GetDiagnosisById(id string) (model.Diagnosis, error)
	GetDiagnosisList() ([]model.Diagnosis, error)
	UpdateDiagnosis(diagnosis model.Diagnosis) (model.Diagnosis, error)
	DeleteDiagnosis(id string) error
}

type Disease interface {
	CreateDisease(disease model.Disease) (model.Disease, error)
	GetDiseaseById(id string) (model.Disease, error)
	GetDiseaseList() ([]model.Disease, error)
	UpdateDisease(disease model.Disease) (model.Disease, error)
	DeleteDisease(id string) error
}

type Doctor interface {
	CreateDoctor(doctor model.Doctor) (model.Doctor, error)
	GetDoctorById(id int) (model.Doctor, error)
	GetDoctorList() ([]model.Doctor, error)
	UpdateDoctor(doctor model.Doctor) (model.Doctor, error)
	DeleteDoctor(id int) error
}

type DoctorPatient interface {
	CreateDoctorPatient(doctorPatient model.DoctorPatient) (model.DoctorPatient, error)
	GetDoctorPatientList(doctor_id int) ([]model.DoctorPatient, error)
	DeleteDoctorPatient(doctor_id, patient_id int) error
}

type Drug interface {
	CreateDrug(drug model.Drug) (model.Drug, error)
	GetDrugById(id string) (model.Drug, error)
	GetDrugList() ([]model.Drug, error)
	UpdateDrug(drug model.Drug) (model.Drug, error)
	DeleteDrug(id string) error
}

type Patient interface {
	CreatePatient(patient model.Patient) (model.Patient, error)
	GetPatientById(id int) (model.Patient, error)
	GetPatientList() ([]model.Patient, error)
	UpdatePatient(patient model.Patient) (model.Patient, error)
	DeletePatient(id int) error
}

type PatientCourse interface {
	CreatePatientCourse(patientCourse model.PatientCourse) (model.PatientCourse, error)
	GetPatientCourseById(id int) (model.PatientCourse, error)
	GetPatientCourseList() ([]model.PatientCourse, error)
	UpdatePatientCourse(patientCourse model.PatientCourse) (model.PatientCourse, error)
	DeletePatientCourse(id int) error
}

type PatientDisease interface {
	CreatePatientDisease(patientDisease model.PatientDisease) (model.PatientDisease, error)
	GetPatientDiseaseById(patientId, diseaseId int) (model.PatientDisease, error)
	GetPatientDiseaseListByPatient(patientId int) ([]model.PatientDisease, error)
	GetPatientDiseaseListByDisease(diseaseId int) ([]model.PatientDisease, error)
	GetPatientDiseaseList() ([]model.PatientDisease, error)
	UpdatePatientDisease(patientDisease model.PatientDisease) (model.PatientDisease, error)
	DeletePatientDisease(patientId, diseaseId int) error
}

type ProcedureBloodCount interface {
	CreateProcedureBloodCount(procedureBloodCount model.ProcedureBloodCount) (model.ProcedureBloodCount, error)
	GetProcedureBloodCountById(procedureId int, bloodCountId string) (model.ProcedureBloodCount, error)
	GetProcedureBloodCountListByProcedure(procedureId int) ([]model.ProcedureBloodCount, error)
	GetProcedureBloodCountListByBloodCount(bloodCountId string) ([]model.ProcedureBloodCount, error)
	GetProcedureBloodCountList() ([]model.ProcedureBloodCount, error)
	UpdateProcedureBloodCount(procedureBloodCount model.ProcedureBloodCount) (model.ProcedureBloodCount, error)
	DeleteProcedureBloodCount(procedureId int, bloodCountId string) error
}

type UnitMeasure interface {
	CreateUnitMeasure(unitMeasure model.UnitMeasure) (model.UnitMeasure, error)
	GetUnitMeasureById(id string) (model.UnitMeasure, error)
	GetUnitMeasureList() ([]model.UnitMeasure, error)
	UpdateUnitMeasure(unitMeasure model.UnitMeasure) (model.UnitMeasure, error)
	DeleteUnitMeasure(id string) error
}

type Service struct {
	Account
	Authorization
	BloodCountValue
	BloodCount
	Course
	CourseProcedure
	Diagnosis
	Disease
	Doctor
	DoctorPatient
	Drug
	Patient
	PatientCourse
	PatientDisease
	ProcedureBloodCount
	UnitMeasure
}

func NewService(repos repository.Repository) *Service {
	return &Service{
		Authorization:       NewAuthService(repos),
		BloodCount:          NewBloodCountService(repos),
		BloodCountValue:     NewBloodCountValueService(repos),
		Course:              NewCourseService(repos),
		CourseProcedure:     NewCourseProcedureService(repos),
		Diagnosis:           NewDiagnosisService(repos),
		Disease:             NewDiseaseService(repos),
		Doctor:              NewDoctorService(repos),
		DoctorPatient:       NewDoctorPatientService(repos),
		Drug:                NewDrugService(repos),
		Patient:             NewPatientService(repos),
		PatientCourse:       NewPatientCourseService(repos),
		PatientDisease:      NewPatientDiseaseService(repos),
		ProcedureBloodCount: NewProcedureBloodCountService(repos),
		UnitMeasure:         NewUnitMeasureService(repos),
	}
}
