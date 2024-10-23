package handler

import (
	"bytes"
	"encoding/json"
	"memrizr/account/model"
	"memrizr/account/model/apperrors"
	"memrizr/account/model/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignup(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Email and Password Required", func(t *testing.T) {
		// we just want this to show that it's not called in this case
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R: router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email": "",
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, req)

		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})

	
	t.Run("Invalid email", func(t *testing.T) {
		// we just want this to show that it's not called in this case
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R: router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email": "bob@bob",
			"password": "supersecret1234",
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, req)

		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})
	
	t.Run("Password too short", func(t *testing.T) {
		// we just want this to show that it's not called in this case
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R: router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email": "bob@bob.com",
			"password": "supe",
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, req)

		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})
	
	t.Run("Password too long", func(t *testing.T) {
		// we just want this to show that it's not called in this case
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R: router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email": "bob@bob.com",
			"password": "vdWYfzJ1PqhMnrKxJyr8cvZ6mPXAyXR3E71No3YBrYktPOCmay",
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, req)

		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})
	
	t.Run("Error returned from UserService", func(t *testing.T) {
		u := &model.User{
			Email: "bob@bob.com",
			Password: "avalidpassword",
		}

		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), u).Return(apperrors.NewConflict("User Already Exist", u.Email))

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R: router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email": u.Email,
			"password": u.Password,
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, req)

		assert.Equal(t, 409, rr.Code)
		mockUserService.AssertExpectations(t)
	})
}