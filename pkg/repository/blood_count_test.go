package repository

import (
	"database/sql"
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

func TestCreateBloodCount(t *testing.T) {
	a := sql.NullString{}
	fmt.Print("a", a)
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewBloodCountRepository(sqlxDB)

	type mockBehavior func(model model.BloodCount)

	testTable := []struct {
		name         string
		model        model.BloodCount
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			model: model.BloodCount{
				Id:               null.StringFrom("1"),
				Description:      null.StringFrom("1"),
				MinNormalValue:   null.FloatFrom(0.1),
				MaxNormalValue:   null.FloatFrom(0.2),
				MinPossibleValue: null.FloatFrom(0.05),
				MaxPossibleValue: null.FloatFrom(0.45),
				MeasureCode:      null.StringFrom("1"),
			},
			mockBehavior: func(model model.BloodCount) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.blood_count").
					WithArgs(
						model.Id,
						model.Description,
						model.MinNormalValue,
						model.MaxNormalValue,
						model.MinPossibleValue,
						model.MaxPossibleValue,
						model.MeasureCode,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Empty required field",
			model: model.BloodCount{
				Description:      null.StringFrom("1"),
				MinNormalValue:   null.FloatFrom(0.1),
				MaxNormalValue:   null.FloatFrom(0.2),
				MinPossibleValue: null.FloatFrom(0.05),
				MaxPossibleValue: null.FloatFrom(0.45),
				MeasureCode:      null.StringFrom("1"),
			},
			mockBehavior: func(model model.BloodCount) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.blood_count").
					WithArgs(
						nil,
						model.Description,
						model.MinNormalValue,
						model.MaxNormalValue,
						model.MinPossibleValue,
						model.MaxPossibleValue,
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
			err := r.CreateBloodCount(testCase.model)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetBloodCountList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewBloodCountRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult []model.BloodCount
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"description",
					"min_normal_value",
					"max_normal_value",
					"min_possible_value",
					"max_possible_value",
					"measure_code",
				}).
					AddRow("1", "descr", 0.1, 0.2, 0.05, 0.45, "1").
					AddRow("2", "", 0.11, 0.21, 0.025, 0.451, "2")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.blood_count").
					WillReturnRows(rows)
			},
			expectResult: []model.BloodCount{
				{
					Id:               null.StringFrom("1"),
					Description:      null.StringFrom("descr"),
					MinNormalValue:   null.FloatFrom(0.1),
					MaxNormalValue:   null.FloatFrom(0.2),
					MinPossibleValue: null.FloatFrom(0.05),
					MaxPossibleValue: null.FloatFrom(0.45),
					MeasureCode:      null.StringFrom("1"),
				},
				{
					Id:               null.StringFrom("2"),
					Description:      null.StringFrom(""),
					MinNormalValue:   null.FloatFrom(0.11),
					MaxNormalValue:   null.FloatFrom(0.21),
					MinPossibleValue: null.FloatFrom(0.025),
					MaxPossibleValue: null.FloatFrom(0.451),
					MeasureCode:      null.StringFrom("2"),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.blood_count").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetBloodCountList()

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetBloodCountById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewBloodCountRepository(sqlxDB)

	type mockBehavior func(id string)

	testTable := []struct {
		name         string
		data         string
		mockBehavior mockBehavior
		expectResult model.BloodCount
		expectErr    bool
	}{
		{
			name: "OK",
			data: "1",
			mockBehavior: func(id string) {
				row := sqlmock.NewRows([]string{
					"id",
					"description",
					"min_normal_value",
					"max_normal_value",
					"min_possible_value",
					"max_possible_value",
					"measure_code",
				}).
					AddRow("1", "descr", 0.1, 0.2, 0.05, 0.45, "1")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.blood_count WHERE id=(.+)").
					WithArgs(id).
					WillReturnRows(row)
			},
			expectResult: model.BloodCount{
				Id:               null.StringFrom("1"),
				Description:      null.StringFrom("descr"),
				MinNormalValue:   null.FloatFrom(0.1),
				MaxNormalValue:   null.FloatFrom(0.2),
				MinPossibleValue: null.FloatFrom(0.05),
				MaxPossibleValue: null.FloatFrom(0.45),
				MeasureCode:      null.StringFrom("1"),
			},
			expectErr: false,
		},
		{
			name: "No records",
			mockBehavior: func(id string) {
				row := sqlmock.NewRows([]string{
					"id",
					"description",
					"min_normal_value",
					"max_normal_value",
					"min_possible_value",
					"max_possible_value",
					"measure_code",
				})

				mock.ExpectQuery("SELECT (.+) FROM onco_base.blood_count WHERE id=(.+)").
					WithArgs(id).
					WillReturnRows(row)
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data)
			res, err := r.GetBloodCountById(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateBloodCount(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewBloodCountRepository(sqlxDB)

	type mockBehavior func(model model.BloodCount)

	testTable := []struct {
		name         string
		data         model.BloodCount
		mockBehavior mockBehavior
		expectResult model.BloodCount
		expectErr    bool
	}{
		{
			name: "OK",
			data: model.BloodCount{
				Id:               null.StringFrom("1"),
				Description:      null.StringFrom("descr"),
				MinNormalValue:   null.FloatFrom(0.1),
				MaxNormalValue:   null.FloatFrom(0.2),
				MinPossibleValue: null.FloatFrom(0.05),
				MaxPossibleValue: null.FloatFrom(0.45),
				MeasureCode:      null.StringFrom("1"),
			},
			mockBehavior: func(model model.BloodCount) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.blood_count SET (.+) WHERE id=(.+)").
					WithArgs(
						model.Description,
						model.MinNormalValue,
						model.MaxNormalValue,
						model.MinPossibleValue,
						model.MaxPossibleValue,
						model.MeasureCode,
						model.Id,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Update error occured",
			mockBehavior: func(model model.BloodCount) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.blood_count SET (.+) WHERE id=(.+)").
					WithArgs(
						model.Description,
						model.MinNormalValue,
						model.MaxNormalValue,
						model.MinPossibleValue,
						model.MaxPossibleValue,
						model.MeasureCode,
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
			err := r.UpdateBloodCount(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDeleteBloodCount(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewBloodCountRepository(sqlxDB)

	type mockBehavior func(id string)

	testTable := []struct {
		name         string
		data         string
		mockBehavior mockBehavior
		expectResult model.BloodCount
		expectErr    bool
	}{
		{
			name: "OK",
			data: "1",
			mockBehavior: func(id string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.blood_count WHERE id=(.+)").
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
				mock.ExpectExec("DELETE FROM onco_base.blood_count WHERE id=(.+)").
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
			err := r.DeleteBloodCount(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
