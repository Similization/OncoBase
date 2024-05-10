package repository

import (
	"errors"
	"log"
	"med/pkg/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
		// expectedBody string
		expectErr bool
	}{
		{
			name: "OK",
			model: model.Disease{
				Id:          "1",
				Description: "text",
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
			name: "Empty Id",
			model: model.Disease{
				Id:          "",
				Description: "text",
			},
			mockBehavior: func(model model.Disease) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.disease").
					WithArgs(model.Id, model.Description).
					WillReturnResult(sqlmock.NewResult(1, 1))
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
					AddRow("1", "description1").
					AddRow("2", "description2").
					AddRow("3", "description3")

				mock.ExpectQuery(`SELECT \* FROM onco_base.disease`).
					WillReturnRows(rows)
			},
			expectResult: []model.Disease{
				{Id: "1", Description: "description1"},
				{Id: "2", Description: "description2"},
				{Id: "3", Description: "description3"},
			},
			expectErr: false,
		},
		{
			name: "Select return error",
			mockBehavior: func() {
				mock.ExpectQuery(`SELECT \* FROM onco_base.disease`).WillReturnError(errors.New("some error occured"))
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
				row := sqlmock.NewRows([]string{"id", "description"}).AddRow("1", "description1")

				mock.ExpectQuery(`SELECT \* FROM onco_base.disease WHERE id=`).WillReturnRows(row)
			},
			expectResult: model.Disease{Id: "1", Description: "description1"},
			expectErr:    false,
		},
		{
			name: "No records",
			mockBehavior: func(id string) {
				row := sqlmock.NewRows([]string{"id", "description"})

				mock.ExpectQuery(`SELECT \* FROM onco_base.disease WHERE id=`).WillReturnRows(row)
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.expectResult.Id)
			res, err := r.GetDiseaseById(testCase.expectResult.Id)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
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
				Id:          "id",
				Description: "new description",
			},
			mockBehavior: func(model.Disease) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.disease SET (.+) WHERE id=(.+)").
					WithArgs("id", "new description").
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
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
				mock.ExpectExec("DELETE FROM onco_base.disease WHERE id=(.+)").WithArgs("1").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Delete return error",
			data: "1",
			mockBehavior: func(id string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.disease WHERE id=(.+)").WithArgs("1").WillReturnError(errors.New("some error occured"))
				mock.ExpectCommit()
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
		})
	}
}
