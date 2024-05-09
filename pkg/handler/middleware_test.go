package handler

import (
	"errors"
	"fmt"
	service "med/pkg/service"
	mock "med/pkg/service/mock"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserIdentity(t *testing.T) {
	type mockBehavior func(s *mock.MockAuthorization, token string)

	testTable := []struct {
		name           string
		headerName     string
		headerValue    string
		token          string
		mockBehavior   mockBehavior
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(&service.UserData{Id: 1}, nil)
			},
			expectedStatus: 200,
			expectedBody:   "1",
		},
		{
			name:           "Invalid Header Name",
			headerName:     "",
			headerValue:    "Bearer token",
			token:          "token",
			mockBehavior:   func(r *mock.MockAuthorization, token string) {},
			expectedStatus: 401,
			expectedBody:   `{"message":"empty auth header"}`,
		},
		{
			name:           "Invalid Header Value",
			headerName:     "Authorization",
			headerValue:    "Bearr token",
			token:          "token",
			mockBehavior:   func(r *mock.MockAuthorization, token string) {},
			expectedStatus: 401,
			expectedBody:   `{"message":"invalid auth header"}`,
		},
		{
			name:           "Empty Token",
			headerName:     "Authorization",
			headerValue:    "Bearer ",
			token:          "token",
			mockBehavior:   func(r *mock.MockAuthorization, token string) {},
			expectedStatus: 401,
			expectedBody:   `{"message":"token is empty"}`,
		},
		{
			name:        "Parse Error",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *mock.MockAuthorization, token string) {
				r.EXPECT().ParseToken(token).Return(&service.UserData{Id: 0}, errors.New("invalid token"))
			},
			expectedStatus: 401,
			expectedBody:   `{"message":"invalid token"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.token)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.GET("/protected", handler.UserIdentity, func(ctx *gin.Context) {
				id, _ := ctx.Get(userContext)
				ctx.String(200, fmt.Sprintf("%d", id))
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatus)
			assert.Equal(t, w.Body.String(), testCase.expectedBody)
		})
	}
}
