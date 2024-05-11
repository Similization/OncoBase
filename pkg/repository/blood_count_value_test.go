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

func TestCreateBloodCountValue(t *testing.T) {
	a := sql.NullString{}
	fmt.Print("a", a)
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewBloodCountValueRepository(sqlxDB)

	type mockBehavior func(model model.BloodCountValue)

	testTable := []struct {
		name         string
		model        model.BloodCountValue
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			model: model.BloodCountValue{
				Disease:     null.StringFrom("1"),
				BloodCount:  null.StringFrom("1"),
				Coefficient: null.FloatFrom(0.1),
				Description: null.StringFrom("descr"),
			},
			mockBehavior: func(model model.BloodCountValue) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.blood_count_value").
					WithArgs(model.Disease, model.BloodCount, model.Coefficient, model.Description).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Empty required field",
			model: model.BloodCountValue{
				BloodCount:  null.StringFrom("1"),
				Coefficient: null.FloatFrom(0.1),
				Description: null.StringFrom("descr"),
			},
			mockBehavior: func(model model.BloodCountValue) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.blood_count_value").
					WithArgs(nil, model.BloodCount, model.Coefficient, model.Description).
					WillReturnError(errors.New("some error occured"))
				mock.ExpectRollback()
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.model)
			err := r.CreateBloodCountValue(testCase.model)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetBloodCountValueList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewBloodCountValueRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult []model.BloodCountValue
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"disease", "blood_count", "coefficient", "description"}).
					AddRow("1", "1", 0.1, "descr").
					AddRow("1", "2", 0.2, "descr2").
					AddRow("1", "3", 0.3, "")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.blood_count_value").
					WillReturnRows(rows)
			},
			expectResult: []model.BloodCountValue{
				{
					Disease:     null.StringFrom("1"),
					BloodCount:  null.StringFrom("1"),
					Coefficient: null.FloatFrom(0.1),
					Description: null.StringFrom("descr"),
				},
				{
					Disease:     null.StringFrom("1"),
					BloodCount:  null.StringFrom("2"),
					Coefficient: null.FloatFrom(0.2),
					Description: null.StringFrom("descr2"),
				},
				{
					Disease:     null.StringFrom("1"),
					BloodCount:  null.StringFrom("3"),
					Coefficient: null.FloatFrom(0.3),
					Description: null.StringFrom(""),
				},
			},
			expectErr: false,
		},
		{
			name: "Select error occured",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.blood_count_value").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetBloodCountValueList()

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetBloodCountValueById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewBloodCountValueRepository(sqlxDB)

	type args struct {
		diseaseId    string
		bloodCountId string
	}

	type mockBehavior func(diseaseId string, bloodCountId string)

	testTable := []struct {
		name         string
		data         args
		mockBehavior mockBehavior
		expectResult model.BloodCountValue
		expectErr    bool
	}{
		{
			name: "OK",
			data: args{
				diseaseId:    "1",
				bloodCountId: "1",
			},
			mockBehavior: func(diseaseId string, bloodCountId string) {
				row := sqlmock.NewRows([]string{"disease", "blood_count", "coefficient", "description"}).
					AddRow("1", "1", 0.1, "descr")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.blood_count_value WHERE disease=(.+) AND blood_count=(.+)").
					WithArgs(diseaseId, bloodCountId).
					WillReturnRows(row)
			},
			expectResult: model.BloodCountValue{
				Disease:     null.StringFrom("1"),
				BloodCount:  null.StringFrom("1"),
				Coefficient: null.FloatFrom(0.1),
				Description: null.StringFrom("descr"),
			},
			expectErr: false,
		},
		{
			name: "No records",
			mockBehavior: func(diseaseId string, bloodCountId string) {
				row := sqlmock.NewRows([]string{"disease", "blood_count", "coefficient", "description"})

				mock.ExpectQuery("SELECT (.+) FROM onco_base.blood_count_value WHERE disease=(.+) AND blood_count=(.+)").
					WithArgs(diseaseId, bloodCountId).
					WillReturnRows(row)
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data.diseaseId, testCase.data.bloodCountId)
			res, err := r.GetBloodCountValueById(testCase.data.diseaseId, testCase.data.bloodCountId)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateBloodCountValue(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewBloodCountValueRepository(sqlxDB)

	type mockBehavior func(model model.BloodCountValue)

	testTable := []struct {
		name         string
		data         model.BloodCountValue
		mockBehavior mockBehavior
		expectResult model.BloodCountValue
		expectErr    bool
	}{
		{
			name: "OK",
			data: model.BloodCountValue{
				Disease:     null.StringFrom("1"),
				BloodCount:  null.StringFrom("1"),
				Coefficient: null.FloatFrom(0.1),
				Description: null.StringFrom("descr"),
			},
			mockBehavior: func(model model.BloodCountValue) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.blood_count_value SET (.+) WHERE disease=(.+) AND blood_count=(.+)").
					WithArgs(model.Coefficient, model.Description, model.Disease, model.BloodCount).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Update error occured",
			mockBehavior: func(model model.BloodCountValue) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.blood_count_value SET (.+) WHERE disease=(.+) AND blood_count=(.+)").
					WithArgs(model.Coefficient, model.Description, model.Disease, model.BloodCount).
					WillReturnError(errors.New("error occured"))
				mock.ExpectRollback()
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data)
			err := r.UpdateBloodCountValue(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDeleteBloodCountValue(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewBloodCountValueRepository(sqlxDB)

	type args struct {
		diseaseId    string
		bloodCountId string
	}

	type mockBehavior func(diseaseId string, bloodCountId string)

	testTable := []struct {
		name         string
		data         args
		mockBehavior mockBehavior
		expectResult model.BloodCountValue
		expectErr    bool
	}{
		{
			name: "OK",
			data: args{
				diseaseId:    "1",
				bloodCountId: "1",
			},
			mockBehavior: func(diseaseId string, bloodCountId string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.blood_count_value WHERE disease=(.+) AND blood_count=(.+)").
					WithArgs(diseaseId, bloodCountId).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Delete error occured",
			data: args{
				diseaseId:    "1",
				bloodCountId: "1",
			},
			mockBehavior: func(diseaseId string, bloodCountId string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.blood_count_value WHERE disease=(.+) AND blood_count=(.+)").
					WithArgs(diseaseId, bloodCountId).
					WillReturnError(errors.New("some error occured"))
				mock.ExpectRollback()
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data.diseaseId, testCase.data.bloodCountId)
			err := r.DeleteBloodCountValue(testCase.data.diseaseId, testCase.data.bloodCountId)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
