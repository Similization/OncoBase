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

func TestCreatePatient(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientRepository(sqlxDB)

	type mockBehavior func(model model.Patient)

	testTable := []struct {
		name         string
		model        model.Patient
		mockBehavior mockBehavior
		expectResult int
		expectErr    bool
	}{
		{
			name: "OK",
			model: model.Patient{
				FirstName:  null.StringFrom("f_name"),
				MiddleName: null.StringFrom("m_name"),
				LastName:   null.StringFrom("l_name"),
				BirthDate:  null.StringFrom("2000-01-01"),
				Sex:        null.StringFrom("m"),
				SNILS:      null.StringFrom("snils"),
				UserId:     null.IntFrom(1),
				Phone:      null.StringFrom("p"),
			},
			mockBehavior: func(model model.Patient) {
				mock.ExpectBegin()

				row := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO onco_base.patient (.+) VALUES (.+) RETURNING id").
					WithArgs(
						model.FirstName,
						model.MiddleName,
						model.LastName,
						model.BirthDate,
						model.Sex,
						model.SNILS,
						model.UserId,
						model.Phone,
					).
					WillReturnRows(row)
				mock.ExpectCommit()
			},
			expectResult: 1,
			expectErr:    false,
		},
		{
			name: "Empty required field",
			model: model.Patient{
				MiddleName: null.StringFrom("m_name"),
				LastName:   null.StringFrom("l_name"),
				BirthDate:  null.StringFrom("2000-01-01"),
				Sex:        null.StringFrom("m"),
				SNILS:      null.StringFrom("snils"),
				UserId:     null.IntFrom(1),
				Phone:      null.StringFrom("p"),
			},
			mockBehavior: func(model model.Patient) {
				mock.ExpectBegin()

				row := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO onco_base.patient (.+) VALUES (.+) RETURNING id").
					WithArgs(
						nil,
						model.MiddleName,
						model.LastName,
						model.BirthDate,
						model.Sex,
						model.SNILS,
						model.UserId,
						model.Phone,
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
			res, err := r.CreatePatient(testCase.model)

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

func TestGetPatientList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult []model.Patient
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"first_name",
					"middle_name",
					"last_name",
					"birth_date",
					"sex",
					"snils",
					"user_id",
					"phone",
				}).
					AddRow(1, "f_name", "m_name", "l_name", "2000-01-01", "m", "snils", 1, "p").
					AddRow(2, "f_name", "", "l_name", "2000-01-02", "f", "snils", 4, "")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient").
					WillReturnRows(rows)
			},
			expectResult: []model.Patient{
				{
					Id:         null.IntFrom(1),
					FirstName:  null.StringFrom("f_name"),
					MiddleName: null.StringFrom("m_name"),
					LastName:   null.StringFrom("l_name"),
					BirthDate:  null.StringFrom("2000-01-01"),
					Sex:        null.StringFrom("m"),
					SNILS:      null.StringFrom("snils"),
					UserId:     null.IntFrom(1),
					Phone:      null.StringFrom("p"),
				},
				{
					Id:         null.IntFrom(2),
					FirstName:  null.StringFrom("f_name"),
					MiddleName: null.StringFrom(""),
					LastName:   null.StringFrom("l_name"),
					BirthDate:  null.StringFrom("2000-01-02"),
					Sex:        null.StringFrom("f"),
					SNILS:      null.StringFrom("snils"),
					UserId:     null.IntFrom(4),
					Phone:      null.StringFrom(""),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetPatientList()

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetPatientById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientRepository(sqlxDB)

	type mockBehavior func(id int)

	testTable := []struct {
		name         string
		data         int
		mockBehavior mockBehavior
		expectResult model.Patient
		expectErr    bool
	}{
		{
			name: "OK",
			data: 1,
			mockBehavior: func(id int) {
				row := sqlmock.NewRows([]string{
					"id",
					"first_name",
					"middle_name",
					"last_name",
					"birth_date",
					"sex",
					"snils",
					"user_id",
					"phone",
				}).
					AddRow(1, "f_name", "m_name", "l_name", "2000-01-01", "m", "snils", 1, "p")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient WHERE id=(.+)").
					WithArgs(id).
					WillReturnRows(row)
			},
			expectResult: model.Patient{
				Id:         null.IntFrom(1),
				FirstName:  null.StringFrom("f_name"),
				MiddleName: null.StringFrom("m_name"),
				LastName:   null.StringFrom("l_name"),
				BirthDate:  null.StringFrom("2000-01-01"),
				Sex:        null.StringFrom("m"),
				SNILS:      null.StringFrom("snils"),
				UserId:     null.IntFrom(1),
				Phone:      null.StringFrom("p"),
			},
			expectErr: false,
		},
		{
			name: "No records",
			mockBehavior: func(id int) {
				row := sqlmock.NewRows([]string{
					"id",
					"first_name",
					"middle_name",
					"last_name",
					"birth_date",
					"sex",
					"snils",
					"user_id",
					"phone",
				})

				mock.ExpectQuery("SELECT (.+) FROM onco_base.patient WHERE id=(.+)").
					WithArgs(id).
					WillReturnRows(row)
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data)
			res, err := r.GetPatientById(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdatePatient(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientRepository(sqlxDB)

	type mockBehavior func(model model.Patient)

	testTable := []struct {
		name         string
		data         model.Patient
		mockBehavior mockBehavior
		expectResult model.Patient
		expectErr    bool
	}{
		{
			name: "OK",
			data: model.Patient{
				Id:         null.IntFrom(1),
				FirstName:  null.StringFrom("f_name"),
				MiddleName: null.StringFrom("m_name"),
				LastName:   null.StringFrom("l_name"),
				BirthDate:  null.StringFrom("2000-01-01"),
				Sex:        null.StringFrom("m"),
				SNILS:      null.StringFrom("snils"),
				UserId:     null.IntFrom(1),
				Phone:      null.StringFrom("p"),
			},
			mockBehavior: func(model model.Patient) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.patient SET (.+) WHERE id = (.+)").
					WithArgs(
						model.FirstName,
						model.MiddleName,
						model.LastName,
						model.BirthDate,
						model.Sex,
						model.SNILS,
						model.UserId,
						model.Phone,
						model.Id,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Update error occured",
			data: model.Patient{
				Id: null.IntFrom(1),
			},
			mockBehavior: func(model model.Patient) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.patient SET (.+) WHERE id = (.+)").
					WithArgs(
						model.FirstName,
						model.MiddleName,
						model.LastName,
						model.BirthDate,
						model.Sex,
						model.SNILS,
						model.UserId,
						model.Phone,
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
			err := r.UpdatePatient(testCase.data)

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

func TestDeletePatient(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewPatientRepository(sqlxDB)

	type mockBehavior func(id int)

	testTable := []struct {
		name         string
		data         int
		mockBehavior mockBehavior
		expectResult model.Patient
		expectErr    bool
	}{
		{
			name: "OK",
			data: 1,
			mockBehavior: func(id int) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.patient WHERE id=(.+)").
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
				mock.ExpectExec("DELETE FROM onco_base.patient WHERE id=(.+)").
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
			err := r.DeletePatient(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
