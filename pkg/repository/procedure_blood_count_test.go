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

func TestCreateProcedureBloodCount(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewProcedureBloodCountRepository(sqlxDB)

	type mockBehavior func(model model.ProcedureBloodCount)

	testTable := []struct {
		name         string
		model        model.ProcedureBloodCount
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			model: model.ProcedureBloodCount{
				Value:       null.FloatFrom(0.1),
				MeasureCode: null.StringFrom("1"),
				Procedure:   null.IntFrom(1),
				BloodCount:  null.StringFrom("1"),
			},
			mockBehavior: func(model model.ProcedureBloodCount) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.procedure_blood_count (.+) VALUES (.+)").
					WithArgs(
						model.Value,
						model.MeasureCode,
						model.Procedure,
						model.BloodCount,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Empty required field",
			model: model.ProcedureBloodCount{
				Value:       null.FloatFrom(0.1),
				MeasureCode: null.StringFrom("1"),
				BloodCount:  null.StringFrom("1"),
			},
			mockBehavior: func(model model.ProcedureBloodCount) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.procedure_blood_count (.+) VALUES (.+)").
					WithArgs(
						model.Value,
						model.MeasureCode,
						nil,
						model.BloodCount,
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
			err := r.CreateProcedureBloodCount(testCase.model)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetProcedureBloodCountList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewProcedureBloodCountRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult []model.ProcedureBloodCount
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{
					"value",
					"measure_code",
					"procedure",
					"blood_count",
				}).
					AddRow(0.1, "1", 1, "1").
					AddRow(0.2, "2", 1, "2")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.procedure_blood_count").
					WillReturnRows(rows)
			},
			expectResult: []model.ProcedureBloodCount{
				{
					Value:       null.FloatFrom(0.1),
					MeasureCode: null.StringFrom("1"),
					Procedure:   null.IntFrom(1),
					BloodCount:  null.StringFrom("1"),
				},
				{
					Value:       null.FloatFrom(0.2),
					MeasureCode: null.StringFrom("2"),
					Procedure:   null.IntFrom(1),
					BloodCount:  null.StringFrom("2"),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.procedure_blood_count").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetProcedureBloodCountList()

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetProcedureBloodCountListByProcedure(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewProcedureBloodCountRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		data         int
		mockBehavior mockBehavior
		expectResult []model.ProcedureBloodCount
		expectErr    bool
	}{
		{
			name: "OK",
			data: 1,
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{
					"value",
					"measure_code",
					"procedure",
					"blood_count",
				}).
					AddRow(0.1, "1", 1, "1").
					AddRow(0.2, "2", 1, "2")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.procedure_blood_count WHERE procedure=(.+)").
					WillReturnRows(rows)
			},
			expectResult: []model.ProcedureBloodCount{
				{
					Value:       null.FloatFrom(0.1),
					MeasureCode: null.StringFrom("1"),
					Procedure:   null.IntFrom(1),
					BloodCount:  null.StringFrom("1"),
				},
				{
					Value:       null.FloatFrom(0.2),
					MeasureCode: null.StringFrom("2"),
					Procedure:   null.IntFrom(1),
					BloodCount:  null.StringFrom("2"),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.procedure_blood_count WHERE procedure=(.+)").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetProcedureBloodCountListByProcedure(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetProcedureBloodCountListByBloodCount(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewProcedureBloodCountRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		data         string
		mockBehavior mockBehavior
		expectResult []model.ProcedureBloodCount
		expectErr    bool
	}{
		{
			name: "OK",
			data: "1",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{
					"value",
					"measure_code",
					"procedure",
					"blood_count",
				}).
					AddRow(0.1, "1", 1, "2").
					AddRow(0.2, "2", 1, "2")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.procedure_blood_count WHERE blood_count=(.+)").
					WillReturnRows(rows)
			},
			expectResult: []model.ProcedureBloodCount{
				{
					Value:       null.FloatFrom(0.1),
					MeasureCode: null.StringFrom("1"),
					Procedure:   null.IntFrom(1),
					BloodCount:  null.StringFrom("2"),
				},
				{
					Value:       null.FloatFrom(0.2),
					MeasureCode: null.StringFrom("2"),
					Procedure:   null.IntFrom(1),
					BloodCount:  null.StringFrom("2"),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.procedure_blood_count WHERE blood_count=(.+)").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetProcedureBloodCountListByBloodCount(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetProcedureBloodCountById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewProcedureBloodCountRepository(sqlxDB)

	type args struct {
		procedureId  int
		bloodCountId string
	}

	type mockBehavior func(procedureId int, bloodCountId string)

	testTable := []struct {
		name         string
		data         args
		mockBehavior mockBehavior
		expectResult model.ProcedureBloodCount
		expectErr    bool
	}{
		{
			name: "OK",
			data: args{
				procedureId:  1,
				bloodCountId: "1",
			},
			mockBehavior: func(procedureId int, bloodCountId string) {
				row := sqlmock.NewRows([]string{
					"value",
					"measure_code",
					"procedure",
					"blood_count",
				}).
					AddRow(0.1, "1", 1, "1")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.procedure_blood_count WHERE procedure=(.+) AND blood_count=(.+)").
					WithArgs(procedureId, bloodCountId).
					WillReturnRows(row)
			},
			expectResult: model.ProcedureBloodCount{
				Value:       null.FloatFrom(0.1),
				MeasureCode: null.StringFrom("1"),
				Procedure:   null.IntFrom(1),
				BloodCount:  null.StringFrom("1"),
			},
			expectErr: false,
		},
		{
			name: "No records",
			mockBehavior: func(procedureId int, bloodCountId string) {
				row := sqlmock.NewRows([]string{
					"value",
					"measure_code",
					"procedure",
					"blood_count",
				})

				mock.ExpectQuery("SELECT (.+) FROM onco_base.procedure_blood_count WHERE procedure=(.+) AND blood_count=(.+)").
					WithArgs(procedureId, bloodCountId).
					WillReturnRows(row)
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data.procedureId, testCase.data.bloodCountId)
			res, err := r.GetProcedureBloodCountById(testCase.data.procedureId, testCase.data.bloodCountId)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateProcedureBloodCount(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewProcedureBloodCountRepository(sqlxDB)

	type mockBehavior func(model model.ProcedureBloodCount)

	testTable := []struct {
		name         string
		data         model.ProcedureBloodCount
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			data: model.ProcedureBloodCount{
				Value:       null.FloatFrom(0.1),
				MeasureCode: null.StringFrom("1"),
				Procedure:   null.IntFrom(1),
				BloodCount:  null.StringFrom("1"),
			},
			mockBehavior: func(model model.ProcedureBloodCount) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.procedure_blood_count SET (.+) WHERE procedure=(.+) AND blood_count=(.+)").
					WithArgs(
						model.Value,
						model.MeasureCode,
						model.Procedure,
						model.BloodCount,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Update error occured",
			mockBehavior: func(model model.ProcedureBloodCount) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.procedure_blood_count SET (.+) WHERE procedure=(.+) AND blood_count=(.+)").
					WithArgs(
						model.Value,
						model.MeasureCode,
						model.Procedure,
						model.BloodCount,
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
			err := r.UpdateProcedureBloodCount(testCase.data)

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

func TestDeleteProcedureBloodCount(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewProcedureBloodCountRepository(sqlxDB)

	type args struct {
		procedureId  int
		bloodCountId string
	}

	type mockBehavior func(procedureId int, bloodCountId string)

	testTable := []struct {
		name         string
		data         args
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			data: args{
				procedureId:  1,
				bloodCountId: "blood_count",
			},
			mockBehavior: func(procedureId int, bloodCountId string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.procedure_blood_count WHERE procedure=(.+) AND blood_count=(.+)").
					WithArgs(procedureId, bloodCountId).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Delete error occured",
			data: args{
				procedureId:  1,
				bloodCountId: "blood_count",
			},
			mockBehavior: func(procedureId int, bloodCountId string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.procedure_blood_count WHERE procedure=(.+) AND blood_count=(.+)").
					WithArgs(procedureId, bloodCountId).
					WillReturnError(errors.New("some error occured"))
				mock.ExpectRollback()
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data.procedureId, testCase.data.bloodCountId)
			err := r.DeleteProcedureBloodCount(testCase.data.procedureId, testCase.data.bloodCountId)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
