package repository

import (
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email, password string) (model.User, error)
}

type Account interface {
}

type BloodCount interface {
	CreateBloodCount(bloodCount model.BloodCount) error
	GetBloodCountById(id string) (model.BloodCount, error)
	GetBloodCountList() ([]model.BloodCount, error)
	UpdateBloodCount(bloodCount model.BloodCount) error
	DeleteBloodCount(id string) error
}

type BloodCountValue interface {
	CreateBloodCountValue(bloodCountValue model.BloodCountValue) error
	GetBloodCountValueById(diseaseId, bloodCountId string) (model.BloodCountValue, error)
	GetBloodCountValueListByDisease(diseaseId string) ([]model.BloodCountValue, error)
	GetBloodCountValueListByBloodCount(bloodCountId string) ([]model.BloodCountValue, error)
	GetBloodCountValueList() ([]model.BloodCountValue, error)
	UpdateBloodCountValue(bloodCountValue model.BloodCountValue) error
	DeleteBloodCountValue(diseaseId, bloodCountId string) error
}

type Course interface {
	CreateCourse(course model.Course) error
	GetCourseById(id string) (model.Course, error)
	GetCourseList() ([]model.Course, error)
	UpdateCourse(course model.Course) error
	DeleteCourse(id string) error
}

type CourseProcedure interface {
	CreateCourseProcedure(courseProcedure model.CourseProcedure) (int, error)
	GetCourseProcedureById(id string) (model.CourseProcedure, error)
	GetCourseProcedureList() ([]model.CourseProcedure, error)
	UpdateCourseProcedure(courseProcedure model.CourseProcedure) error
	DeleteCourseProcedure(id string) error
}

type Diagnosis interface {
	CreateDiagnosis(diagnosis model.Diagnosis) error
	GetDiagnosisById(id string) (model.Diagnosis, error)
	GetDiagnosisList() ([]model.Diagnosis, error)
	UpdateDiagnosis(diagnosis model.Diagnosis) error
	DeleteDiagnosis(id string) error
}

type Disease interface {
	CreateDisease(disease model.Disease) error
	GetDiseaseById(id string) (model.Disease, error)
	GetDiseaseList() ([]model.Disease, error)
	UpdateDisease(disease model.Disease) error
	DeleteDisease(id string) error
}

type Doctor interface {
	CreateDoctor(doctor model.Doctor) (int, error)
	GetDoctorById(id int) (model.Doctor, error)
	GetDoctorList() ([]model.Doctor, error)
	UpdateDoctor(doctor model.Doctor) error
	DeleteDoctor(id int) error
}

type DoctorPatient interface {
	CreateDoctorPatient(doctorPatient model.DoctorPatient) error
	GetDoctorPatientListByDoctor(doctorId int) ([]model.DoctorPatient, error)
	GetDoctorPatientListByPatient(patientId int) ([]model.DoctorPatient, error)
	DeleteDoctorPatient(doctorPatient model.DoctorPatient) error
}

type Drug interface {
	CreateDrug(drug model.Drug) error
	GetDrugById(id string) (model.Drug, error)
	GetDrugList() ([]model.Drug, error)
	UpdateDrug(drug model.Drug) error
	DeleteDrug(id string) error
}

type Patient interface {
	CreatePatient(patient model.Patient) (int, error)
	GetPatientById(id int) (model.Patient, error)
	GetPatientList() ([]model.Patient, error)
	UpdatePatient(patient model.Patient) error
	DeletePatient(id int) error
}

type PatientCourse interface {
	CreatePatientCourse(patientCourse model.PatientCourse) (int, error)
	GetPatientCourseById(id int) (model.PatientCourse, error)
	GetPatientCourseList() ([]model.PatientCourse, error)
	UpdatePatientCourse(patientCourse model.PatientCourse) error
	DeletePatientCourse(id int) error
}

type PatientDisease interface {
	CreatePatientDisease(patientDisease model.PatientDisease) error
	GetPatientDiseaseById(patientId int, diseaseId string) (model.PatientDisease, error)
	GetPatientDiseaseListByPatient(patientId int) ([]model.PatientDisease, error)
	GetPatientDiseaseListByDisease(diseaseId string) ([]model.PatientDisease, error)
	GetPatientDiseaseList() ([]model.PatientDisease, error)
	UpdatePatientDisease(patientDisease model.PatientDisease) error
	DeletePatientDisease(patientId int, diseaseId string) error
}

type ProcedureBloodCount interface {
	CreateProcedureBloodCount(procedureBloodCount model.ProcedureBloodCount) error
	GetProcedureBloodCountById(procedureId int, bloodCountId string) (model.ProcedureBloodCount, error)
	GetProcedureBloodCountListByProcedure(procedureId int) ([]model.ProcedureBloodCount, error)
	GetProcedureBloodCountListByBloodCount(bloodCountId string) ([]model.ProcedureBloodCount, error)
	GetProcedureBloodCountList() ([]model.ProcedureBloodCount, error)
	UpdateProcedureBloodCount(procedureBloodCount model.ProcedureBloodCount) error
	DeleteProcedureBloodCount(procedureId int, bloodCountId string) error
}

type UnitMeasure interface {
	CreateUnitMeasure(unitMeasure model.UnitMeasure) error
	GetUnitMeasureById(id string) (model.UnitMeasure, error)
	GetUnitMeasureList() ([]model.UnitMeasure, error)
	UpdateUnitMeasure(unitMeasure model.UnitMeasure) error
	DeleteUnitMeasure(id string) error
}

type Repository struct {
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:       NewAuthRepository(db),
		BloodCount:          NewBloodCountRepository(db),
		BloodCountValue:     NewBloodCountValueRepository(db),
		Course:              NewCourseRepository(db),
		CourseProcedure:     NewCourseProcedureRepository(db),
		Diagnosis:           NewDiagnosisRepository(db),
		Disease:             NewDiseaseRepository(db),
		Doctor:              NewDoctorRepository(db),
		DoctorPatient:       NewDoctorPatientRepository(db),
		Drug:                NewDrugRepository(db),
		Patient:             NewPatientRepository(db),
		PatientCourse:       NewPatientCourseRepository(db),
		PatientDisease:      NewPatientDiseaseRepository(db),
		ProcedureBloodCount: NewProcedureBloodCountRepository(db),
		UnitMeasure:         NewUnitMeasureRepository(db),
	}
}
