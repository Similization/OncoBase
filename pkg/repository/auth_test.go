package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"med/pkg/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	a := sql.NullString{}
	fmt.Print("a", a)
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewAuthRepository(sqlxDB)

	type mockBehavior func(model model.User)

	testTable := []struct {
		name         string
		model        model.User
		mockBehavior mockBehavior
		expectResult int
		expectErr    bool
	}{
		{
			name: "OK",
			model: model.User{
				Email:    "email",
				Password: "pass",
				Role:     "role",
			},
			mockBehavior: func(model model.User) {
				mock.ExpectBegin()

				row := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO onco_base.external_user").
					WithArgs(model.Email, model.Password, model.Role).
					WillReturnRows(row)
				mock.ExpectCommit()
			},
			expectResult: 1,
			expectErr:    false,
		},
		{
			name: "Empty required field",
			model: model.User{
				Password: "pass",
				Role:     "role",
			},
			mockBehavior: func(model model.User) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO onco_base.external_user").
					WithArgs(nil, model.Password, model.Role).
					WillReturnError(errors.New("some error occured"))
				mock.ExpectRollback()
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.model)
			res, err := r.CreateUser(testCase.model)

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

func TestGetUser(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewAuthRepository(sqlxDB)

	type args struct {
		email    string
		password string
	}

	type mockBehavior func(email string, password string)

	testTable := []struct {
		name         string
		data         args
		mockBehavior mockBehavior
		expectResult model.User
		expectErr    bool
	}{
		{
			name: "OK",
			mockBehavior: func(email string, password string) {
				rows := sqlmock.NewRows([]string{"id", "description"}).
					AddRow(1, "email", "pass", "role")

				mock.ExpectQuery("SELECT (.+) FROM onco_base.internal_user WHERE email=(.+) AND password=(.+)").
					WillReturnRows(rows)
			},
			expectResult: model.User{Id: 1, Email: "email", Password: "pass", Role: "role"},
			expectErr:    false,
		},
		{
			name: "Select error occured",
			mockBehavior: func(email string, password string) {
				mock.ExpectQuery("SELECT (.+) FROM onco_base.external_user").
					WillReturnError(errors.New("some error occured"))
			},
			expectErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.data.email, testCase.data.password)
			res, err := r.GetUser(testCase.data.email, testCase.data.password)

			if testCase.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectResult, res)
				assert.NoError(t, err)
			}
		})
	}
}
