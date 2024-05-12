package repository

import (
	"errors"
	"fmt"
	"log"
	"med/pkg/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/guregu/null/v5"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreatePatientDisease(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientDiseaseRepository(sqlxDB)

	type mockBehavior func(model model.PatientDisease)

	testTable := []struct {
		name         string
		model        model.PatientDisease
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			model: model.PatientDisease{
				Patient:   null.IntFrom(1),
				Disease:   null.StringFrom("disease"),
				Stage:     null.StringFrom("stage"),
				Diagnosis: null.StringFrom("diagnosis"),
			},
			mockBehavior: func(model model.PatientDisease) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.patient_disease (.+) VALUES (.+)").
					WithArgs(
						model.Stage,
						model.Diagnosis,
						model.Patient,
						model.Disease,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Empty required field",
			model: model.PatientDisease{
				Disease:   null.StringFrom("disease"),
				Stage:     null.StringFrom("stage"),
				Diagnosis: null.StringFrom("diagnosis"),
			},
			mockBehavior: func(model model.PatientDisease) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.patient_disease (.+) VALUES (.+)").
					WithArgs(
						model.Stage,
						model.Diagnosis,
						nil,
						model.Disease,
					).
					WillReturnError(errors.New("empty required field"))
				mock.ExpectRollback()
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.model)
			err := r.CreatePatientDisease(testCase.model)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetPatientDiseaseList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientDiseaseRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult []model.PatientDisease
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{
					"patient",
					"disease",
					"stage",
					"diagnosis",
				}).
					AddRow(1, "disease", "stage", "diagnosis").
					AddRow(2, "disease", "", "diagnosis")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient_disease").
					WillReturnRows(rows)
			},
			expectResult: []model.PatientDisease{
				{
					Patient:   null.IntFrom(1),
					Disease:   null.StringFrom("disease"),
					Stage:     null.StringFrom("stage"),
					Diagnosis: null.StringFrom("diagnosis"),
				},
				{
					Patient:   null.IntFrom(2),
					Disease:   null.StringFrom("disease"),
					Stage:     null.StringFrom(""),
					Diagnosis: null.StringFrom("diagnosis"),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient_disease").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetPatientDiseaseList()

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetPatientDiseaseListByPatient(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientDiseaseRepository(sqlxDB)

	type mockBehavior func(patientId int)

	testTable := []struct {
		name         string
		data         int
		mockBehavior mockBehavior
		expectResult []model.PatientDisease
		expectErr    bool
	}{
		{
			name: "OK",
			data: 1,
			mockBehavior: func(patientId int) {
				rows := sqlmock.NewRows([]string{
					"patient",
					"disease",
					"stage",
					"diagnosis",
				}).
					AddRow(1, "disease", "stage", "diagnosis").
					AddRow(1, "disease2", "", "diagnosis")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient_disease WHERE patient=(.+)").
					WithArgs(patientId).
					WillReturnRows(rows)
			},
			expectResult: []model.PatientDisease{
				{
					Patient:   null.IntFrom(1),
					Disease:   null.StringFrom("disease"),
					Stage:     null.StringFrom("stage"),
					Diagnosis: null.StringFrom("diagnosis"),
				},
				{
					Patient:   null.IntFrom(1),
					Disease:   null.StringFrom("disease2"),
					Stage:     null.StringFrom(""),
					Diagnosis: null.StringFrom("diagnosis"),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func(patientId int) {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient_disease WHERE patient=(.+)").
					WithArgs(patientId).
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data)
			res, err := r.GetPatientDiseaseListByPatient(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetPatientDiseaseListByDisease(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientDiseaseRepository(sqlxDB)

	type mockBehavior func(diseaseId string)

	testTable := []struct {
		name         string
		data         string
		mockBehavior mockBehavior
		expectResult []model.PatientDisease
		expectErr    bool
	}{
		{
			name: "OK",
			data: "disease",
			mockBehavior: func(diseaseId string) {
				rows := sqlmock.NewRows([]string{
					"patient",
					"disease",
					"stage",
					"diagnosis",
				}).
					AddRow(1, "disease", "stage", "diagnosis").
					AddRow(2, "disease", "", "diagnosis")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient_disease WHERE disease=(.+)").
					WithArgs(diseaseId).
					WillReturnRows(rows)
			},
			expectResult: []model.PatientDisease{
				{
					Patient:   null.IntFrom(1),
					Disease:   null.StringFrom("disease"),
					Stage:     null.StringFrom("stage"),
					Diagnosis: null.StringFrom("diagnosis"),
				},
				{
					Patient:   null.IntFrom(2),
					Disease:   null.StringFrom("disease"),
					Stage:     null.StringFrom(""),
					Diagnosis: null.StringFrom("diagnosis"),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func(diseaseId string) {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient_disease WHERE disease=(.+)").
					WithArgs(diseaseId).
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data)
			res, err := r.GetPatientDiseaseListByDisease(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetPatientDiseaseById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientDiseaseRepository(sqlxDB)

	type args struct {
		patientId int
		diseaseId string
	}

	type mockBehavior func(patientId int, diseaseId string)

	testTable := []struct {
		name         string
		data         args
		mockBehavior mockBehavior
		expectResult model.PatientDisease
		expectErr    bool
	}{
		{
			name: "OK",
			data: args{
				patientId: 1,
				diseaseId: "1",
			},
			mockBehavior: func(patientId int, diseaseId string) {
				row := sqlmock.NewRows([]string{
					"patient",
					"disease",
					"stage",
					"diagnosis",
				}).
					AddRow(1, "1", "stage", "diagnosis")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient_disease WHERE patient=(.+) AND disease=(.+)").
					WithArgs(patientId, diseaseId).
					WillReturnRows(row)
			},
			expectResult: model.PatientDisease{
				Patient:   null.IntFrom(1),
				Disease:   null.StringFrom("1"),
				Stage:     null.StringFrom("stage"),
				Diagnosis: null.StringFrom("diagnosis"),
			},
			expectErr: false,
		},
		{
			name: "No records",
			mockBehavior: func(patientId int, diseaseId string) {
				row := sqlmock.NewRows([]string{
					"patient",
					"disease",
					"stage",
					"diagnosis",
				})

				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient_disease WHERE patient=(.+) AND disease=(.+)").
					WithArgs(patientId, diseaseId).
					WillReturnRows(row)
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data.patientId, testCase.data.diseaseId)
			res, err := r.GetPatientDiseaseById(testCase.data.patientId, testCase.data.diseaseId)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdatePatientDisease(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientDiseaseRepository(sqlxDB)

	type mockBehavior func(model model.PatientDisease)

	testTable := []struct {
		name         string
		data         model.PatientDisease
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			data: model.PatientDisease{
				Patient:   null.IntFrom(1),
				Disease:   null.StringFrom("disease"),
				Stage:     null.StringFrom("stage"),
				Diagnosis: null.StringFrom("diagnosis"),
			},
			mockBehavior: func(model model.PatientDisease) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.patient_disease SET (.+) WHERE patient=(.+) AND disease=(.+)").
					WithArgs(
						model.Stage,
						model.Diagnosis,
						model.Patient,
						model.Disease,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Update error occured",
			mockBehavior: func(model model.PatientDisease) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.patient_disease SET (.+) WHERE patient=(.+) AND disease=(.+)").
					WithArgs(
						model.Patient,
						model.Disease,
						model.Stage,
						model.Diagnosis,
					).
					WillReturnError(errors.New("error occured"))
				mock.ExpectRollback()
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data)
			err := r.UpdatePatientDisease(testCase.data)

			if testCase.expectErr {
				fmt.Print(err)
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDeletePatientDisease(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientDiseaseRepository(sqlxDB)

	type args struct {
		patientId int
		diseaseId string
	}

	type mockBehavior func(patientId int, diseaseId string)

	testTable := []struct {
		name         string
		data         args
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			data: args{
				patientId: 1,
				diseaseId: "disease",
			},
			mockBehavior: func(patientId int, diseaseId string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.patient_disease WHERE patient=(.+) AND disease=(.+)").
					WithArgs(patientId, diseaseId).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Delete error occured",
			data: args{
				patientId: 1,
				diseaseId: "disease",
			},
			mockBehavior: func(patientId int, diseaseId string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.patient_disease WHERE patient=(.+) AND disease=(.+)").
					WithArgs(patientId, diseaseId).
					WillReturnError(errors.New("some error occured"))
				mock.ExpectRollback()
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data.patientId, testCase.data.diseaseId)
			err := r.DeletePatientDisease(testCase.data.patientId, testCase.data.diseaseId)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
