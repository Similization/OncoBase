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

func TestCreateDiagnosis(t *testing.T) {
	a := sql.NullString{}
	fmt.Print("a", a)
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDiagnosisRepository(sqlxDB)

	type mockBehavior func(model model.Diagnosis)

	testTable := []struct {
		name         string
		model        model.Diagnosis
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			model: model.Diagnosis{
				Id:          null.StringFrom("1"),
				Description: null.StringFrom("text"),
			},
			mockBehavior: func(model model.Diagnosis) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.diagnosis").
					WithArgs(model.Id, model.Description).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Empty required field",
			model: model.Diagnosis{
				Description: null.StringFrom("text"),
			},
			mockBehavior: func(model model.Diagnosis) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.diagnosis").
					WithArgs(nil, model.Description).
					WillReturnError(errors.New("some error occured"))
				mock.ExpectRollback()
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.model)
			err := r.CreateDiagnosis(testCase.model)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetDiagnosisList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDiagnosisRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult []model.Diagnosis
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id", "description"}).
					AddRow("1", "descr1").
					AddRow("2", "descr2").
					AddRow("3", "descr3")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.diagnosis").
					WillReturnRows(rows)
			},
			expectResult: []model.Diagnosis{
				{
					Id:          null.StringFrom("1"),
					Description: null.StringFrom("descr1"),
				},
				{
					Id:          null.StringFrom("2"),
					Description: null.StringFrom("descr2"),
				},
				{
					Id:          null.StringFrom("3"),
					Description: null.StringFrom("descr3"),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.diagnosis").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetDiagnosisList()

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetDiagnosisById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDiagnosisRepository(sqlxDB)

	type mockBehavior func(id string)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult model.Diagnosis
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func(id string) {
				row := sqlmock.NewRows([]string{"id", "description"}).AddRow("1", "description1")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.diagnosis WHERE id=").WithArgs(id).WillReturnRows(row)
			},
			expectResult: model.Diagnosis{
				Id:          null.StringFrom("1"),
				Description: null.StringFrom("description1"),
			},
			expectErr: false,
		},
		{
			name: "No records",
			mockBehavior: func(id string) {
				row := sqlmock.NewRows([]string{"id", "description"})

				mock.ExpectQuery("SELECT (.+) FROM onco_base.diagnosis WHERE id=(.+)").WithArgs(id).WillReturnRows(row)
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.expectResult.Id.String)
			res, err := r.GetDiagnosisById(testCase.expectResult.Id.String)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateDiagnosis(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDiagnosisRepository(sqlxDB)

	type mockBehavior func(model model.Diagnosis)

	testTable := []struct {
		name         string
		data         model.Diagnosis
		mockBehavior mockBehavior
		expectResult model.Diagnosis
		expectErr    bool
	}{
		{
			name: "OK",
			data: model.Diagnosis{
				Id:          null.StringFrom("1"),
				Description: null.StringFrom("description1"),
			},
			mockBehavior: func(model model.Diagnosis) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.diagnosis SET (.+) WHERE id=(.+)").
					WithArgs(model.Id, model.Description).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Update error occured",
			mockBehavior: func(model model.Diagnosis) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.diagnosis SET (.+) WHERE id=(.+)").
					WithArgs(model.Id, model.Description).
					WillReturnError(errors.New("error occured"))
				mock.ExpectRollback()
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data)
			err := r.UpdateDiagnosis(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDeleteDiagnosis(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDiagnosisRepository(sqlxDB)

	type mockBehavior func(id string)

	testTable := []struct {
		name         string
		data         string
		mockBehavior mockBehavior
		expectResult model.Diagnosis
		expectErr    bool
	}{
		{
			name: "OK",
			data: "1",
			mockBehavior: func(id string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.diagnosis WHERE id=(.+)").
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
				mock.ExpectExec("DELETE FROM onco_base.diagnosis WHERE id=(.+)").
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
			err := r.DeleteDiagnosis(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
