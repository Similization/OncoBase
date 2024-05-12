package repository

import (
	"errors"
	"log"
	"med/pkg/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/guregu/null/v5"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateCourseProcedure(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewCourseProcedureRepository(sqlxDB)

	type mockBehavior func(model model.CourseProcedure)

	testTable := []struct {
		name         string
		model        model.CourseProcedure
		mockBehavior mockBehavior
		expectResult int
		expectErr    bool
	}{
		{
			name: "OK",
			model: model.CourseProcedure{
				PatientCourse: null.IntFrom(1),
				Doctor:        null.IntFrom(1),
				BeginDate:     null.StringFrom(time.Now().String()),
				Period:        null.IntFrom(7),
				Result:        null.StringFrom("res"),
			},
			mockBehavior: func(model model.CourseProcedure) {
				mock.ExpectBegin()
				row := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO onco_base.course_procedure (.+) VALUES (.+) RETURNING id").
					WithArgs(
						model.PatientCourse,
						model.Doctor,
						model.BeginDate,
						model.Period,
						model.Result,
					).
					WillReturnRows(row)
				mock.ExpectCommit()
			},
			expectResult: 1,
			expectErr:    false,
		},
		{
			name: "Empty required field",
			model: model.CourseProcedure{
				Doctor:    null.IntFrom(1),
				BeginDate: null.StringFrom(time.Now().String()),
				Period:    null.IntFrom(7),
				Result:    null.StringFrom("res"),
			},
			mockBehavior: func(model model.CourseProcedure) {
				mock.ExpectBegin()
				row := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO onco_base.course_procedure (.+) VALUES (.+) RETURNING id").
					WithArgs(
						nil,
						model.Doctor,
						model.BeginDate,
						model.Period,
						model.Result,
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
			res, err := r.CreateCourseProcedure(testCase.model)

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

func TestGetCourseProcedureList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewCourseProcedureRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult []model.CourseProcedure
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"patient_course",
					"doctor",
					"begin_date",
					"period",
					"result",
				}).
					AddRow(1, 1, 1, "2020-01-01", 7, "res").
					AddRow(2, 1, 2, "2020-01-02", 14, "")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.course_procedure").
					WillReturnRows(rows)
			},
			expectResult: []model.CourseProcedure{
				{
					Id:            null.IntFrom(1),
					PatientCourse: null.IntFrom(1),
					Doctor:        null.IntFrom(1),
					BeginDate:     null.StringFrom("2020-01-01"),
					Period:        null.IntFrom(7),
					Result:        null.StringFrom("res"),
				},
				{
					Id:            null.IntFrom(2),
					PatientCourse: null.IntFrom(1),
					Doctor:        null.IntFrom(2),
					BeginDate:     null.StringFrom("2020-01-02"),
					Period:        null.IntFrom(14),
					Result:        null.StringFrom(""),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.course_procedure").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetCourseProcedureList()

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetCourseProcedureById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewCourseProcedureRepository(sqlxDB)

	type mockBehavior func(id string)

	testTable := []struct {
		name         string
		data         string
		mockBehavior mockBehavior
		expectResult model.CourseProcedure
		expectErr    bool
	}{
		{
			name: "OK",
			data: "1",
			mockBehavior: func(id string) {
				row := sqlmock.NewRows([]string{
					"id",
					"patient_course",
					"doctor",
					"begin_date",
					"period",
					"result",
				}).
					AddRow(1, 1, 1, "2020-01-01", 7, "res")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.course_procedure WHERE id=(.+)").
					WithArgs(id).
					WillReturnRows(row)
			},
			expectResult: model.CourseProcedure{
				Id:            null.IntFrom(1),
				PatientCourse: null.IntFrom(1),
				Doctor:        null.IntFrom(1),
				BeginDate:     null.StringFrom("2020-01-01"),
				Period:        null.IntFrom(7),
				Result:        null.StringFrom("res"),
			},
			expectErr: false,
		},
		{
			name: "No records",
			mockBehavior: func(id string) {
				row := sqlmock.NewRows([]string{
					"id",
					"period",
					"frequency",
					"dose",
					"drug",
					"measure_code",
				})

				mock.ExpectQuery("SELECT (.+) FROM onco_base.course_procedure WHERE id=(.+)").
					WithArgs(id).
					WillReturnRows(row)
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data)
			res, err := r.GetCourseProcedureById(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateCourseProcedure(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewCourseProcedureRepository(sqlxDB)

	type mockBehavior func(model model.CourseProcedure)

	testTable := []struct {
		name         string
		data         model.CourseProcedure
		mockBehavior mockBehavior
		expectResult model.CourseProcedure
		expectErr    bool
	}{
		{
			name: "OK",
			data: model.CourseProcedure{
				Id:            null.IntFrom(1),
				PatientCourse: null.IntFrom(1),
				Doctor:        null.IntFrom(1),
				BeginDate:     null.StringFrom("2020-01-01"),
				Period:        null.IntFrom(7),
				Result:        null.StringFrom("res"),
			},
			mockBehavior: func(model model.CourseProcedure) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.course_procedure SET (.+) WHERE id=(.+)").
					WithArgs(
						model.PatientCourse,
						model.Doctor,
						model.BeginDate,
						model.Period,
						model.Result,
						model.Id,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Update error occured",
			mockBehavior: func(model model.CourseProcedure) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.course_procedure SET (.+) WHERE id=(.+)").
					WithArgs(
						model.PatientCourse,
						model.Doctor,
						model.BeginDate,
						model.Period,
						model.Result,
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
			err := r.UpdateCourseProcedure(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDeleteCourseProcedure(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewCourseProcedureRepository(sqlxDB)

	type mockBehavior func(id string)

	testTable := []struct {
		name         string
		data         string
		mockBehavior mockBehavior
		expectResult model.CourseProcedure
		expectErr    bool
	}{
		{
			name: "OK",
			data: "1",
			mockBehavior: func(id string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.course_procedure WHERE id=(.+)").
					WithArgs(id).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Delete error occured",
			data: "1",
			mockBehavior: func(id string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.course_procedure WHERE id=(.+)").
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
			err := r.DeleteCourseProcedure(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
