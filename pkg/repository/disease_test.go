package repository

import (
	"database/sql"
	"errors"
	"log"
	"med/pkg/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/guregu/null/v5"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateDisease(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDiseaseRepository(sqlxDB)

	type mockBehavior func(model model.Disease)

	testTable := []struct {
		name         string
		model        model.Disease
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			model: model.Disease{
				Id:          null.String{NullString: sql.NullString{String: "1", Valid: true}},
				Description: null.String{NullString: sql.NullString{String: "desc", Valid: true}},
			},
			mockBehavior: func(model model.Disease) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.disease").
					WithArgs(model.Id, model.Description).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Empty required field",
			model: model.Disease{
				Description: null.String{NullString: sql.NullString{String: "descr", Valid: true}},
			},
			mockBehavior: func(model model.Disease) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.disease").
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
			err := r.CreateDisease(testCase.model)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetDiseaseList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDiseaseRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult []model.Disease
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id", "description"}).
					AddRow("1", "descr1").
					AddRow("2", "descr2").
					AddRow("3", "descr3")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.disease").
					WillReturnRows(rows)
			},
			expectResult: []model.Disease{
				{
					Id:          null.String{NullString: sql.NullString{String: "1", Valid: true}},
					Description: null.String{NullString: sql.NullString{String: "descr1", Valid: true}},
				},
				{
					Id:          null.String{NullString: sql.NullString{String: "2", Valid: true}},
					Description: null.String{NullString: sql.NullString{String: "descr2", Valid: true}},
				},
				{
					Id:          null.String{NullString: sql.NullString{String: "3", Valid: true}},
					Description: null.String{NullString: sql.NullString{String: "descr3", Valid: true}},
				},
			},
			expectErr: false,
		},
		{
			name: "Select return error",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.disease").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetDiseaseList()

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

func TestGetDiseaseById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDiseaseRepository(sqlxDB)

	type mockBehavior func(id string)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult model.Disease
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func(id string) {
				row := sqlmock.NewRows([]string{"id", "description"}).AddRow("1", "descr")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.disease WHERE id=").WithArgs(id).WillReturnRows(row)
			},
			expectResult: model.Disease{
				Id:          null.String{NullString: sql.NullString{String: "1", Valid: true}},
				Description: null.String{NullString: sql.NullString{String: "descr", Valid: true}},
			},
			expectErr: false,
		},
		{
			name: "No records",
			mockBehavior: func(id string) {
				row := sqlmock.NewRows([]string{"id", "description"})

				mock.ExpectQuery("SELECT (.+) FROM onco_base.disease WHERE id=").WithArgs(id).WillReturnRows(row)
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.expectResult.Id.String)
			res, err := r.GetDiseaseById(testCase.expectResult.Id.String)

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

func TestUpdateDisease(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDiseaseRepository(sqlxDB)

	type mockBehavior func(model model.Disease)

	testTable := []struct {
		name         string
		data         model.Disease
		mockBehavior mockBehavior
		expectResult model.Disease
		expectErr    bool
	}{
		{
			name: "OK",
			data: model.Disease{
				Id:          null.String{NullString: sql.NullString{String: "1", Valid: true}},
				Description: null.String{NullString: sql.NullString{String: "new descr", Valid: true}},
			},
			mockBehavior: func(model model.Disease) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.disease SET (.+) WHERE id=(.+)").
					WithArgs(model.Id, model.Description).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Update error occured",
			mockBehavior: func(model model.Disease) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.disease SET (.+) WHERE id=(.+)").
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
			err := r.UpdateDisease(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDeleteDisease(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDiseaseRepository(sqlxDB)

	type mockBehavior func(id string)

	testTable := []struct {
		name         string
		data         string
		mockBehavior mockBehavior
		expectResult model.Disease
		expectErr    bool
	}{
		{
			name: "OK",
			data: "1",
			mockBehavior: func(id string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.disease WHERE id=(.+)").
					WithArgs(id).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Delete return error",
			data: "1",
			mockBehavior: func(id string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.disease WHERE id=(.+)").
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
			err := r.DeleteDisease(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
