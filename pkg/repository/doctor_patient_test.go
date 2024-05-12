package repository

import (
	"errors"
	"log"
	"med/pkg/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/guregu/null/v5"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateDoctorPatient(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDoctorPatientRepository(sqlxDB)

	type mockBehavior func(model model.DoctorPatient)

	testTable := []struct {
		name         string
		model        model.DoctorPatient
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			model: model.DoctorPatient{
				Patient: null.IntFrom(1),
				Doctor:  null.IntFrom(1),
			},
			mockBehavior: func(model model.DoctorPatient) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.doctor_patient (.+) VALUES (.+)").
					WithArgs(
						model.Patient,
						model.Doctor,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Empty required field",
			model: model.DoctorPatient{
				Doctor: null.IntFrom(1),
			},
			mockBehavior: func(model model.DoctorPatient) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.doctor_patient (.+) VALUES (.+)").
					WithArgs(
						nil,
						model.Doctor,
					).
					WillReturnError(errors.New("Empty required field"))
				mock.ExpectRollback()
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.model)
			err := r.CreateDoctorPatient(testCase.model)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetDoctorPatientListByDoctor(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDoctorPatientRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		data         int
		mockBehavior mockBehavior
		expectResult []model.DoctorPatient
		expectErr    bool
	}{
		{
			name: "OK",
			data: 1,
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{
					"patient",
					"doctor",
				}).
					AddRow(1, 1).
					AddRow(2, 1)

				mock.ExpectQuery("SELECT (.+) FROM onco_base.doctor_patient WHERE doctor=(.+)").
					WillReturnRows(rows)
			},
			expectResult: []model.DoctorPatient{
				{
					Patient: null.IntFrom(1),
					Doctor:  null.IntFrom(1),
				},
				{
					Patient: null.IntFrom(2),
					Doctor:  null.IntFrom(1),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.doctor_patient WHERE doctor=(.+)").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetDoctorPatientListByDoctor(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetDoctorPatientListByPatient(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDoctorPatientRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		data         int
		mockBehavior mockBehavior
		expectResult []model.DoctorPatient
		expectErr    bool
	}{
		{
			name: "OK",
			data: 1,
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{
					"patient",
					"doctor",
				}).
					AddRow(1, 1).
					AddRow(1, 2)

				mock.ExpectQuery("SELECT (.+) FROM onco_base.doctor_patient WHERE patient=(.+)").
					WillReturnRows(rows)
			},
			expectResult: []model.DoctorPatient{
				{
					Patient: null.IntFrom(1),
					Doctor:  null.IntFrom(1),
				},
				{
					Patient: null.IntFrom(1),
					Doctor:  null.IntFrom(2),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.doctor_patient WHERE patient=(.+)").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetDoctorPatientListByPatient(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteDoctorPatient(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDoctorPatientRepository(sqlxDB)

	type mockBehavior func(data model.DoctorPatient)

	testTable := []struct {
		name         string
		data         model.DoctorPatient
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			data: model.DoctorPatient{
				Patient: null.IntFrom(1),
				Doctor:  null.IntFrom(1),
			},
			mockBehavior: func(data model.DoctorPatient) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.doctor_patient WHERE patient=(.+) AND doctor=(.+)").
					WithArgs(data.Patient, data.Doctor).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Delete error occured",
			data: model.DoctorPatient{
				Patient: null.IntFrom(1),
				Doctor:  null.IntFrom(1),
			},
			mockBehavior: func(data model.DoctorPatient) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.doctor_patient WHERE patient=(.+) AND doctor=(.+)").
					WithArgs(data.Patient, data.Doctor).
					WillReturnError(errors.New("some error occured"))
				mock.ExpectRollback()
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data)
			err := r.DeleteDoctorPatient(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
