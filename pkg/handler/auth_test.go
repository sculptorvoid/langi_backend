package handler

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/sculptorvoid/langi_backend/pkg/entity"
	"github.com/sculptorvoid/langi_backend/pkg/service"
	service_mocks "github.com/sculptorvoid/langi_backend/pkg/service/mocks"
	"net/http/httptest"
	"testing"
)

func TestHandler_registration(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *service_mocks.MockAuthorization, user entity.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            entity.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Correct values, will be success",
			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty"}`,
			inputUser: entity.User{
				Username: "username",
				Name:     "Test Name",
				Password: "qwerty",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user entity.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"username": "username"}`,
			inputUser:            entity.User{},
			mockBehavior:         func(r *service_mocks.MockAuthorization, user entity.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty"}`,
			inputUser: entity.User{
				Username: "username",
				Name:     "Test Name",
				Password: "qwerty",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user entity.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/registration", handler.registration)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/registration",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
