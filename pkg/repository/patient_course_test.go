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

func TestCreatePatientCourse(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientCourseRepository(sqlxDB)

	type mockBehavior func(model model.PatientCourse)

	testTable := []struct {
		name         string
		model        model.PatientCourse
		mockBehavior mockBehavior
		expectResult int
		expectErr    bool
	}{
		{
			name: "OK",
			model: model.PatientCourse{
				Patient:   null.IntFrom(1),
				Disease:   null.StringFrom("disease"),
				Course:    null.StringFrom("course"),
				Doctor:    null.IntFrom(1),
				BeginDate: null.StringFrom("2000-01-01"),
				EndDate:   null.StringFrom("2000-01-11"),
				Diagnosis: null.StringFrom("diagnosis"),
			},
			mockBehavior: func(model model.PatientCourse) {
				mock.ExpectBegin()

				row := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO onco_base.patient_course (.+) VALUES (.+) RETURNING id").
					WithArgs(
						model.Patient,
						model.Disease,
						model.Course,
						model.Doctor,
						model.BeginDate,
						model.EndDate,
						model.Diagnosis,
					).
					WillReturnRows(row)
				mock.ExpectCommit()
			},
			expectResult: 1,
			expectErr:    false,
		},
		{
			name: "Empty required field",
			model: model.PatientCourse{
				Disease:   null.StringFrom("disease"),
				Course:    null.StringFrom("course"),
				Doctor:    null.IntFrom(1),
				BeginDate: null.StringFrom("2000-01-01"),
				EndDate:   null.StringFrom("2000-01-11"),
				Diagnosis: null.StringFrom("diagnosis"),
			},
			mockBehavior: func(model model.PatientCourse) {
				mock.ExpectBegin()

				row := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO onco_base.patient_course (.+) VALUES (.+) RETURNING id").
					WithArgs(
						nil,
						model.Disease,
						model.Course,
						model.Doctor,
						model.BeginDate,
						model.EndDate,
						model.Diagnosis,
					).
					WillReturnRows(row)
				mock.ExpectRollback()
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.model)
			res, err := r.CreatePatientCourse(testCase.model)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetPatientCourseList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientCourseRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult []model.PatientCourse
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"patient",
					"disease",
					"course",
					"doctor",
					"begin_date",
					"end_date",
					"diagnosis",
				}).
					AddRow(1, 1, "disease", "course", 1, "2000-01-01", "2000-01-11", "diagnosis").
					AddRow(2, 4, "disease", "course", 3, "2000-01-02", "", "diagnosis")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient_course").
					WillReturnRows(rows)
			},
			expectResult: []model.PatientCourse{
				{
					Id:        null.IntFrom(1),
					Patient:   null.IntFrom(1),
					Disease:   null.StringFrom("disease"),
					Course:    null.StringFrom("course"),
					Doctor:    null.IntFrom(1),
					BeginDate: null.StringFrom("2000-01-01"),
					EndDate:   null.StringFrom("2000-01-11"),
					Diagnosis: null.StringFrom("diagnosis"),
				},
				{
					Id:        null.IntFrom(2),
					Patient:   null.IntFrom(4),
					Disease:   null.StringFrom("disease"),
					Course:    null.StringFrom("course"),
					Doctor:    null.IntFrom(3),
					BeginDate: null.StringFrom("2000-01-02"),
					EndDate:   null.StringFrom(""),
					Diagnosis: null.StringFrom("diagnosis"),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient_course").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetPatientCourseList()

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetPatientCourseById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientCourseRepository(sqlxDB)

	type mockBehavior func(id int)

	testTable := []struct {
		name         string
		data         int
		mockBehavior mockBehavior
		expectResult model.PatientCourse
		expectErr    bool
	}{
		{
			name: "OK",
			data: 1,
			mockBehavior: func(id int) {
				row := sqlmock.NewRows([]string{
					"id",
					"patient",
					"disease",
					"course",
					"doctor",
					"begin_date",
					"end_date",
					"diagnosis",
				}).
					AddRow(1, 1, "disease", "course", 1, "2000-01-01", "2000-01-11", "diagnosis")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient_course WHERE id=(.+)").
					WithArgs(id).
					WillReturnRows(row)
			},
			expectResult: model.PatientCourse{
				Id:        null.IntFrom(1),
				Patient:   null.IntFrom(1),
				Disease:   null.StringFrom("disease"),
				Course:    null.StringFrom("course"),
				Doctor:    null.IntFrom(1),
				BeginDate: null.StringFrom("2000-01-01"),
				EndDate:   null.StringFrom("2000-01-11"),
				Diagnosis: null.StringFrom("diagnosis"),
			},
			expectErr: false,
		},
		{
			name: "No records",
			mockBehavior: func(id int) {
				row := sqlmock.NewRows([]string{
					"id",
					"patient",
					"disease",
					"course",
					"doctor",
					"begin_date",
					"end_date",
					"diagnosis",
				})

				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient_course WHERE id=(.+)").
					WithArgs(id).
					WillReturnRows(row)
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data)
			res, err := r.GetPatientCourseById(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdatePatientCourse(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientCourseRepository(sqlxDB)

	type mockBehavior func(model model.PatientCourse)

	testTable := []struct {
		name         string
		data         model.PatientCourse
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			data: model.PatientCourse{
				Id:        null.IntFrom(1),
				Patient:   null.IntFrom(1),
				Disease:   null.StringFrom("disease"),
				Course:    null.StringFrom("course"),
				Doctor:    null.IntFrom(1),
				BeginDate: null.StringFrom("2000-01-01"),
				EndDate:   null.StringFrom("2000-01-11"),
				Diagnosis: null.StringFrom("diagnosis"),
			},
			mockBehavior: func(model model.PatientCourse) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.patient_course SET (.+) WHERE id=(.+)").
					WithArgs(
						model.Patient,
						model.Disease,
						model.Course,
						model.Doctor,
						model.BeginDate,
						model.EndDate,
						model.Diagnosis,
						model.Id,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Update error occured",
			data: model.PatientCourse{
				Id: null.IntFrom(1),
			},
			mockBehavior: func(model model.PatientCourse) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.patient_course SET (.+) WHERE id=(.+)").
					WithArgs(
						model.Patient,
						model.Disease,
						model.Course,
						model.Doctor,
						model.BeginDate,
						model.EndDate,
						model.Diagnosis,
						model.Id,
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
			err := r.UpdatePatientCourse(testCase.data)

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

func TestDeletePatientCourse(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientCourseRepository(sqlxDB)

	type mockBehavior func(id int)

	testTable := []struct {
		name         string
		data         int
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			data: 1,
			mockBehavior: func(id int) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.patient_course WHERE id=(.+)").
					WithArgs(id).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Delete error occured",
			data: 1,
			mockBehavior: func(id int) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.patient_course WHERE id=(.+)").
					WithArgs(id).
					WillReturnError(errors.New("some error occured"))
				mock.ExpectRollback()
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data)
			err := r.DeletePatientCourse(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
