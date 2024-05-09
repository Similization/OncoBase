package handler

import (
	"bytes"
	"errors"
	model "med/pkg/model"
	service "med/pkg/service"
	mock "med/pkg/service/mock"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRegistry(t *testing.T) {

	type mockBehavior func(s *mock.MockAuthorization, user model.User)

	testTable := []struct {
		name           string
		inputBody      string
		inputUser      model.User
		mockBehavior   mockBehavior
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "OK",
			inputBody: `{"email": "user_email", "password": "pass", "role": "doctor"}`,
			inputUser: model.User{
				Email:    "user_email",
				Password: "pass",
				Role:     "doctor",
			},
			mockBehavior: func(s *mock.MockAuthorization, user model.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatus: 200,
			expectedBody:   `{"id":1}`,
		},
		{
			name:           "Empty fields",
			inputBody:      `{"password": "pass", "role": "doctor"}`,
			mockBehavior:   func(s *mock.MockAuthorization, user model.User) {},
			expectedStatus: 400,
			expectedBody:   `{"message":"Invalid input body"}`,
		},
		{
			name:      "Service error",
			inputBody: `{"email": "user_email", "password": "pass", "role": "doctor"}`,
			inputUser: model.User{
				Email:    "user_email",
				Password: "pass",
				Role:     "doctor",
			},
			mockBehavior: func(s *mock.MockAuthorization, user model.User) {
				s.EXPECT().CreateUser(user).Return(0, errors.New("Internal server error"))
			},
			expectedStatus: 500,
			expectedBody:   `{"message":"Internal server error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/registry", handler.Registry)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/registry", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatus, w.Code)
			assert.Equal(t, testCase.expectedBody, w.Body.String())
		})
	}
}
