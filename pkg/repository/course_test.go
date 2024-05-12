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

func TestCreateCourse(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewCourseRepository(sqlxDB)

	type mockBehavior func(model model.Course)

	testTable := []struct {
		name         string
		model        model.Course
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			model: model.Course{
				Id:          null.StringFrom("1"),
				Period:      null.IntFrom(1),
				Frequency:   null.FloatFrom(0.1),
				Dose:        null.FloatFrom(0.2),
				Drug:        null.StringFrom("1"),
				MeasureCode: null.StringFrom("1"),
			},
			mockBehavior: func(model model.Course) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.course").
					WithArgs(
						model.Id,
						model.Period,
						model.Frequency,
						model.Dose,
						model.Drug,
						model.MeasureCode,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Empty required field",
			model: model.Course{
				Period:      null.IntFrom(1),
				Frequency:   null.FloatFrom(0.1),
				Dose:        null.FloatFrom(0.2),
				Drug:        null.StringFrom("1"),
				MeasureCode: null.StringFrom("1"),
			},
			mockBehavior: func(model model.Course) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.course").
					WithArgs(
						nil,
						model.Period,
						model.Frequency,
						model.Dose,
						model.Drug,
						model.MeasureCode,
					).
					WillReturnError(errors.New("some error occured"))
				mock.ExpectRollback()
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.model)
			err := r.CreateCourse(testCase.model)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetCourseList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewCourseRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult []model.Course
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"period",
					"frequency",
					"dose",
					"drug",
					"measure_code",
				}).
					AddRow("1", 1, 0.1, 0.2, "1", "1").
					AddRow("2", 2, 0.2, 0.3, "2", "1")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.course").
					WillReturnRows(rows)
			},
			expectResult: []model.Course{
				{
					Id:          null.StringFrom("1"),
					Period:      null.IntFrom(1),
					Frequency:   null.FloatFrom(0.1),
					Dose:        null.FloatFrom(0.2),
					Drug:        null.StringFrom("1"),
					MeasureCode: null.StringFrom("1"),
				},
				{
					Id:          null.StringFrom("2"),
					Period:      null.IntFrom(2),
					Frequency:   null.FloatFrom(0.2),
					Dose:        null.FloatFrom(0.3),
					Drug:        null.StringFrom("2"),
					MeasureCode: null.StringFrom("1"),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.course").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetCourseList()

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetCourseById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewCourseRepository(sqlxDB)

	type mockBehavior func(id string)

	testTable := []struct {
		name         string
		data         string
		mockBehavior mockBehavior
		expectResult model.Course
		expectErr    bool
	}{
		{
			name: "OK",
			data: "1",
			mockBehavior: func(id string) {
				row := sqlmock.NewRows([]string{
					"id",
					"period",
					"frequency",
					"dose",
					"drug",
					"measure_code",
				}).
					AddRow("1", 1, 0.1, 0.2, "1", "1")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.course WHERE id=(.+)").
					WithArgs(id).
					WillReturnRows(row)
			},
			expectResult: model.Course{
				Id:          null.StringFrom("1"),
				Period:      null.IntFrom(1),
				Frequency:   null.FloatFrom(0.1),
				Dose:        null.FloatFrom(0.2),
				Drug:        null.StringFrom("1"),
				MeasureCode: null.StringFrom("1"),
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

				mock.ExpectQuery("SELECT (.+) FROM onco_base.course WHERE id=(.+)").
					WithArgs(id).
					WillReturnRows(row)
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data)
			res, err := r.GetCourseById(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateCourse(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewCourseRepository(sqlxDB)

	type mockBehavior func(model model.Course)

	testTable := []struct {
		name         string
		data         model.Course
		mockBehavior mockBehavior
		expectResult model.Course
		expectErr    bool
	}{
		{
			name: "OK",
			data: model.Course{
				Id:          null.StringFrom("1"),
				Period:      null.IntFrom(1),
				Frequency:   null.FloatFrom(0.1),
				Dose:        null.FloatFrom(0.2),
				Drug:        null.StringFrom("1"),
				MeasureCode: null.StringFrom("1"),
			},
			mockBehavior: func(model model.Course) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.course SET (.+) WHERE id=(.+)").
					WithArgs(
						model.Id,
						model.Period,
						model.Frequency,
						model.Dose,
						model.Drug,
						model.MeasureCode,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Update error occured",
			mockBehavior: func(model model.Course) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.course SET (.+) WHERE id=(.+)").
					WithArgs(
						model.Id,
						model.Period,
						model.Frequency,
						model.Dose,
						model.Drug,
						model.MeasureCode,
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
			err := r.UpdateCourse(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDeleteCourse(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewCourseRepository(sqlxDB)

	type mockBehavior func(id string)

	testTable := []struct {
		name         string
		data         string
		mockBehavior mockBehavior
		expectResult model.Course
		expectErr    bool
	}{
		{
			name: "OK",
			data: "1",
			mockBehavior: func(id string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.course WHERE id=(.+)").
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
				mock.ExpectExec("DELETE FROM onco_base.course WHERE id=(.+)").
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
			err := r.DeleteCourse(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
