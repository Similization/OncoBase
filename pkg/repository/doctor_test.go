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

func TestCreateDoctor(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDoctorRepository(sqlxDB)

	type mockBehavior func(model model.Doctor)

	testTable := []struct {
		name         string
		model        model.Doctor
		mockBehavior mockBehavior
		expectResult int
		expectErr    bool
	}{
		{
			name: "OK",
			model: model.Doctor{
				FirstName:     null.StringFrom("f_name"),
				MiddleName:    null.StringFrom("m_name"),
				LastName:      null.StringFrom("l_name"),
				Qualification: null.StringFrom("q"),
				Phone:         null.StringFrom("p"),
				UserId:        null.IntFrom(1),
			},
			mockBehavior: func(model model.Doctor) {
				mock.ExpectBegin()

				row := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO onco_base.doctor (.+) VALUES (.+) RETURNING id").
					WithArgs(
						model.FirstName,
						model.MiddleName,
						model.LastName,
						model.Qualification,
						model.Phone,
						model.UserId,
					).
					WillReturnRows(row)
				mock.ExpectCommit()
			},
			expectResult: 1,
			expectErr:    false,
		},
		{
			name: "Empty required field",
			model: model.Doctor{
				MiddleName:    null.StringFrom("m_name"),
				LastName:      null.StringFrom("l_name"),
				Qualification: null.StringFrom("q"),
				Phone:         null.StringFrom("p"),
				UserId:        null.IntFrom(1),
			},
			mockBehavior: func(model model.Doctor) {
				mock.ExpectBegin()

				row := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO onco_base.doctor (.+) VALUES (.+) RETURNING id").
					WithArgs(
						nil,
						model.MiddleName,
						model.LastName,
						model.Qualification,
						model.Phone,
						model.UserId,
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
			res, err := r.CreateDoctor(testCase.model)

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

func TestGetDoctorList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDoctorRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult []model.Doctor
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
					"qualification",
					"phone",
					"user_id",
				}).
					AddRow(1, "f_name", "m_name", "l_name", "q", "p", 1).
					AddRow(2, "f_name", "m_name", "l_name", "q", "", 3)

				mock.ExpectQuery("SELECT (.+) FROM onco_base.doctor").
					WillReturnRows(rows)
			},
			expectResult: []model.Doctor{
				{
					Id:            null.IntFrom(1),
					FirstName:     null.StringFrom("f_name"),
					MiddleName:    null.StringFrom("m_name"),
					LastName:      null.StringFrom("l_name"),
					Qualification: null.StringFrom("q"),
					Phone:         null.StringFrom("p"),
					UserId:        null.IntFrom(1),
				},
				{
					Id:            null.IntFrom(2),
					FirstName:     null.StringFrom("f_name"),
					MiddleName:    null.StringFrom("m_name"),
					LastName:      null.StringFrom("l_name"),
					Qualification: null.StringFrom("q"),
					Phone:         null.StringFrom(""),
					UserId:        null.IntFrom(3),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.doctor").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetDoctorList()

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetDoctorById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDoctorRepository(sqlxDB)

	type mockBehavior func(id int)

	testTable := []struct {
		name         string
		data         int
		mockBehavior mockBehavior
		expectResult model.Doctor
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
					"qualification",
					"phone",
					"user_id",
				}).
					AddRow(1, "f_name", "m_name", "l_name", "q", "p", 1)

				mock.ExpectQuery("SELECT (.+) FROM onco_base.doctor WHERE id=(.+)").
					WithArgs(id).
					WillReturnRows(row)
			},
			expectResult: model.Doctor{
				Id:            null.IntFrom(1),
				FirstName:     null.StringFrom("f_name"),
				MiddleName:    null.StringFrom("m_name"),
				LastName:      null.StringFrom("l_name"),
				Qualification: null.StringFrom("q"),
				Phone:         null.StringFrom("p"),
				UserId:        null.IntFrom(1),
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
					"qualification",
					"phone",
					"user_id",
				})

				mock.ExpectQuery("SELECT (.+) FROM onco_base.doctor WHERE id=(.+)").
					WithArgs(id).
					WillReturnRows(row)
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data)
			res, err := r.GetDoctorById(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateDoctor(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDoctorRepository(sqlxDB)

	type mockBehavior func(model model.Doctor)

	testTable := []struct {
		name         string
		data         model.Doctor
		mockBehavior mockBehavior
		expectResult model.Doctor
		expectErr    bool
	}{
		{
			name: "OK",
			data: model.Doctor{
				Id:            null.IntFrom(1),
				FirstName:     null.StringFrom("f_name"),
				MiddleName:    null.StringFrom("m_name"),
				LastName:      null.StringFrom("l_name"),
				Qualification: null.StringFrom("q"),
				Phone:         null.StringFrom("p"),
				UserId:        null.IntFrom(1),
			},
			mockBehavior: func(model model.Doctor) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.doctor SET (.+) WHERE id=(.+)").
					WithArgs(
						model.FirstName,
						model.MiddleName,
						model.LastName,
						model.Qualification,
						model.Phone,
						model.UserId,
						model.Id,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Update error occured",
			mockBehavior: func(model model.Doctor) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.doctor SET (.+) WHERE id=(.+)").
					WithArgs(
						model.FirstName,
						model.MiddleName,
						model.LastName,
						model.Qualification,
						model.Phone,
						model.UserId,
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
			err := r.UpdateDoctor(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDeleteDoctor(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDoctorRepository(sqlxDB)

	type mockBehavior func(id int)

	testTable := []struct {
		name         string
		data         int
		mockBehavior mockBehavior
		expectResult model.Doctor
		expectErr    bool
	}{
		{
			name: "OK",
			data: 1,
			mockBehavior: func(id int) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.doctor WHERE id=(.+)").
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
				mock.ExpectExec("DELETE FROM onco_base.doctor WHERE id=(.+)").
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
			err := r.DeleteDoctor(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
