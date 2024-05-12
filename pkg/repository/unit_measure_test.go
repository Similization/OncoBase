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

func TestCreateUnitMeasure(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewUnitMeasureRepository(sqlxDB)

	type mockBehavior func(model model.UnitMeasure)

	testTable := []struct {
		name         string
		model        model.UnitMeasure
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			model: model.UnitMeasure{
				Id:        null.StringFrom("1"),
				Shorthand: null.StringFrom("shrt"),
				FullText:  null.StringFrom("full text"),
				Global:    null.StringFrom("glob"),
			},
			mockBehavior: func(model model.UnitMeasure) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.unit_measure (.+) VALUES (.+)").
					WithArgs(
						model.Id,
						model.Shorthand,
						model.FullText,
						model.Global,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Empty required field",
			model: model.UnitMeasure{
				Shorthand: null.StringFrom("shrt"),
				FullText:  null.StringFrom("full text"),
				Global:    null.StringFrom("glob"),
			},
			mockBehavior: func(model model.UnitMeasure) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.unit_measure (.+) VALUES (.+)").
					WithArgs(
						nil,
						model.Shorthand,
						model.FullText,
						model.Global,
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
			err := r.CreateUnitMeasure(testCase.model)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetUnitMeasureList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewUnitMeasureRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult []model.UnitMeasure
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"shorthand",
					"full_text",
					"global",
				}).
					AddRow("1", "shrt", "full text", "glob").
					AddRow("2", "shrt", "full text", "glob")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.unit_measure").
					WillReturnRows(rows)
			},
			expectResult: []model.UnitMeasure{
				{
					Id:        null.StringFrom("1"),
					Shorthand: null.StringFrom("shrt"),
					FullText:  null.StringFrom("full text"),
					Global:    null.StringFrom("glob"),
				},
				{
					Id:        null.StringFrom("2"),
					Shorthand: null.StringFrom("shrt"),
					FullText:  null.StringFrom("full text"),
					Global:    null.StringFrom("glob"),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.unit_measure").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetUnitMeasureList()

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetUnitMeasureById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewUnitMeasureRepository(sqlxDB)

	type mockBehavior func(id string)

	testTable := []struct {
		name         string
		data         string
		mockBehavior mockBehavior
		expectResult model.UnitMeasure
		expectErr    bool
	}{
		{
			name: "OK",
			data: "1",
			mockBehavior: func(id string) {
				row := sqlmock.NewRows([]string{
					"id",
					"shorthand",
					"full_text",
					"global",
				}).
					AddRow("1", "shrt", "full text", "glob")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.unit_measure WHERE id=(.+)").
					WithArgs(id).
					WillReturnRows(row)
			},
			expectResult: model.UnitMeasure{
				Id:        null.StringFrom("1"),
				Shorthand: null.StringFrom("shrt"),
				FullText:  null.StringFrom("full text"),
				Global:    null.StringFrom("glob"),
			},
			expectErr: false,
		},
		{
			name: "No records",
			mockBehavior: func(id string) {
				row := sqlmock.NewRows([]string{
					"value",
					"measure_code",
					"procedure",
					"blood_count",
				})

				mock.ExpectQuery("SELECT (.+) FROM onco_base.unit_measure WHERE id=(.+)").
					WithArgs(id).
					WillReturnRows(row)
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data)
			res, err := r.GetUnitMeasureById(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateUnitMeasure(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewUnitMeasureRepository(sqlxDB)

	type mockBehavior func(model model.UnitMeasure)

	testTable := []struct {
		name         string
		data         model.UnitMeasure
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			data: model.UnitMeasure{
				Id:        null.StringFrom("1"),
				Shorthand: null.StringFrom("shrt"),
				FullText:  null.StringFrom("full text"),
				Global:    null.StringFrom("glob"),
			},
			mockBehavior: func(model model.UnitMeasure) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.unit_measure SET (.+) WHERE id=(.+)").
					WithArgs(
						model.Id,
						model.Shorthand,
						model.FullText,
						model.Global,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Update error occured",
			mockBehavior: func(model model.UnitMeasure) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.unit_measure SET (.+) WHERE id=(.+)").
					WithArgs(
						model.Id,
						model.Shorthand,
						model.FullText,
						model.Global,
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
			err := r.UpdateUnitMeasure(testCase.data)

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

func TestDeleteUnitMeasure(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewUnitMeasureRepository(sqlxDB)

	type mockBehavior func(id string)

	testTable := []struct {
		name         string
		data         string
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			data: "1",
			mockBehavior: func(id string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.unit_measure WHERE id=(.+)").
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
				mock.ExpectExec("DELETE FROM onco_base.unit_measure WHERE id=(.+)").
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
			err := r.DeleteUnitMeasure(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
