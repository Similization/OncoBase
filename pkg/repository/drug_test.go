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

func TestCreateDrug(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDrugRepository(sqlxDB)

	type mockBehavior func(model model.Drug)

	testTable := []struct {
		name         string
		model        model.Drug
		mockBehavior mockBehavior
		expectErr    bool
	}{
		{
			name: "OK",
			model: model.Drug{
				Id:                "1",
				Name:              "name",
				DosageForm:        "form",
				ActiveIngredients: "ingridients",
				Country:           "cou",
				Manufacturer:      "man",
				PrescribingOrder:  "ord",
				Description:       "text",
			},
			mockBehavior: func(model model.Drug) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.drug").
					WithArgs(
						model.Id,
						model.Name,
						model.DosageForm,
						model.ActiveIngredients,
						model.Country,
						model.Manufacturer,
						model.PrescribingOrder,
						model.Description,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Empty required field",
			mockBehavior: func(model model.Drug) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.drug").
					WithArgs(
						nil,
						model.Name,
						model.DosageForm,
						model.ActiveIngredients,
						model.Country,
						model.Manufacturer,
						model.PrescribingOrder,
						model.Description,
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
			err := r.CreateDrug(testCase.model)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetDrugList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDrugRepository(sqlxDB)

	type mockBehavior func()

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult []model.Drug
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{
					"id", "name", "dosage_form", "active_ingredients", "country", "manufacturer", "prescribing_order", "description",
				}).
					AddRow("1", "name", "form", "ingridients", "cou", "man", "ord", "text").
					AddRow("2", "name2", "form2", "ingridients2", "", "", "", "text2")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.drug").
					WillReturnRows(rows)
			},
			expectResult: []model.Drug{
				{
					Id:                "1",
					Name:              "name",
					DosageForm:        "form",
					ActiveIngredients: "ingridients",
					Country:           "cou",
					Manufacturer:      "man",
					PrescribingOrder:  "ord",
					Description:       "text",
				},
				{
					Id:                "2",
					Name:              "name2",
					DosageForm:        "form2",
					ActiveIngredients: "ingridients2",
					Description:       "text2",
				},
			},
			expectErr: false,
		},
		{
			name: "Select return error",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.drug").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			res, err := r.GetDrugList()

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

func TestGetDrugById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDrugRepository(sqlxDB)

	type mockBehavior func(id string)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectResult model.Drug
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func(id string) {
				row := sqlmock.NewRows([]string{
					"id", "name", "dosage_form", "active_ingredients", "country", "manufacturer", "prescribing_order", "description",
				}).
					AddRow("1", "name", "form", "ingridients", "cou", "man", "ord", "text")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.drug WHERE id=(.+)").WithArgs(id).WillReturnRows(row)
			},
			expectResult: model.Drug{
				Id:                "1",
				Name:              "name",
				DosageForm:        "form",
				ActiveIngredients: "ingridients",
				Country:           "cou",
				Manufacturer:      "man",
				PrescribingOrder:  "ord",
				Description:       "text",
			},
			expectErr: false,
		},
		{
			name: "No records",
			mockBehavior: func(id string) {
				row := sqlmock.NewRows([]string{"id", "description"})

				mock.ExpectQuery("SELECT (.+) FROM onco_base.drug WHERE id=(.+)").WillReturnRows(row)
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.expectResult.Id)
			res, err := r.GetDrugById(testCase.expectResult.Id)

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

func TestUpdateDrug(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDrugRepository(sqlxDB)

	type mockBehavior func(model model.Drug)

	testTable := []struct {
		name         string
		data         model.Drug
		mockBehavior mockBehavior
		expectResult model.Drug
		expectErr    bool
	}{
		{
			name: "OK",
			data: model.Drug{
				Id:                "1",
				Name:              "name",
				DosageForm:        "form",
				ActiveIngredients: "ingridients",
				Country:           "cou",
				Manufacturer:      "man",
				PrescribingOrder:  "ord",
				Description:       "text",
			},
			mockBehavior: func(model model.Drug) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.drug SET (.+) WHERE id=(.+)").
					WithArgs(
						model.Id,
						model.Name,
						model.DosageForm,
						model.ActiveIngredients,
						model.Country,
						model.Manufacturer,
						model.PrescribingOrder,
						model.Description,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectErr: false,
		},
		{
			name: "Update error occured",
			mockBehavior: func(model model.Drug) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE onco_base.drug SET (.+) WHERE id=(.+)").
					WithArgs(
						model.Id,
						model.Name,
						model.DosageForm,
						model.ActiveIngredients,
						model.Country,
						model.Manufacturer,
						model.PrescribingOrder,
						model.Description,
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
			err := r.UpdateDrug(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDeleteDrug(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewDrugRepository(sqlxDB)

	type mockBehavior func(id string)

	testTable := []struct {
		name         string
		data         string
		mockBehavior mockBehavior
		expectResult model.Drug
		expectErr    bool
	}{
		{
			name: "OK",
			data: "1",
			mockBehavior: func(id string) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM onco_base.drug WHERE id=(.+)").
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
				mock.ExpectExec("DELETE FROM onco_base.drug WHERE id=(.+)").
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
			err := r.DeleteDrug(testCase.data)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
