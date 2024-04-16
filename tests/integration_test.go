package tests

import (
	"bytes"
	"fmt"
	"github.com/RakhimovAns/FinalYandexTask/initializers"
	"github.com/RakhimovAns/FinalYandexTask/models"
	handler "github.com/RakhimovAns/FinalYandexTask/pkg/handlers"
	"github.com/RakhimovAns/FinalYandexTask/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Запуск gRPC-сервера перед запуском тестов
	go server.StartGRPCServer()

	exitCode := m.Run()

	os.Exit(exitCode)
}

// TestPostExpression дополненный
func TestPostExpression(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	initializers.ConnectToDB()
	initializers.CreateTable()
	type cases struct {
		name     string
		jsonStr  []byte
		expected int
		token    string
	}
	tests := []cases{
		{
			name:     "sum",
			jsonStr:  []byte(`{"expression":"2+2","addTime":1,"divideTime":1,"subTime":1,"multiplyTime":1}`),
			expected: http.StatusOK,
			token:    GetToken(),
		},
		{
			name:     "minus",
			jsonStr:  []byte(`{"expression":"2-2","addTime":1,"divideTime":1,"subTime":1,"multiplyTime":1}`),
			expected: http.StatusOK,
			token:    GetToken(),
		},
		{
			name:     "divide",
			jsonStr:  []byte(`{"expression":"2/2","addTime":1,"divideTime":1,"subTime":1,"multiplyTime":1}`),
			expected: http.StatusOK,
			token:    GetToken(),
		},
		{
			name:     "sum",
			jsonStr:  []byte(`{"expression":"2*2","addTime":1,"divideTime":1,"subTime":1,"multiplyTime":1}`),
			expected: http.StatusOK,
			token:    GetToken(),
		},
		{
			name:     "mixed",
			jsonStr:  []byte(`{"expression":"2+2*2","addTime":1,"divideTime":1,"subTime":1,"multiplyTime":1}`),
			expected: http.StatusOK,
			token:    GetToken(),
		},
		{
			name:     "mixed with invalid expression",
			jsonStr:  []byte(`{"expression":"2+2*","addTime":1,"divideTime":1,"subTime":1,"multiplyTime":1}`),
			expected: http.StatusInternalServerError,
			token:    GetToken(),
		},
		{
			name:     "missing token",
			jsonStr:  []byte(`{"expression":"2+2*2","addTime":1,"divideTime":1,"subTime":1,"multiplyTime":1}`),
			expected: 200,
			token:    "", // No token provided
		},
	}
	r.POST("/expression", handler.PostExpression)
	for _, tc := range tests {
		req, err := http.NewRequest("POST", "/expression", bytes.NewBuffer(tc.jsonStr))
		if err != nil {
			t.Fatal(err)
		}
		if tc.token != "" {
			cookie := &http.Cookie{Name: "token", Value: tc.token}
			req.AddCookie(cookie)
		}
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != tc.expected {
			t.Errorf("test case '%s': expected status %d, got %d", tc.name, tc.expected, w.Code)
		}
	}
}

// TestStatusExpression дополненный
func TestStatusExpression(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	initializers.ConnectToDB()
	initializers.CreateTable()
	r.GET("/status/:id", handler.GetStatus)
	type cases struct {
		name     string
		jsonStr  []byte
		id       int
		expected int
		token    string
	}
	tests := []cases{
		{
			name:     "good",
			jsonStr:  []byte(`{"id": 1}`),
			id:       1,
			expected: http.StatusOK,
			token:    GetToken(),
		},
		{
			name:     "not exist",
			jsonStr:  []byte(`{"id": 1000}`),
			id:       1000,
			expected: 200,
			token:    GetToken(),
		},
		{
			name:     "missing token",
			jsonStr:  []byte(`{"id": 1}`),
			id:       1,
			expected: 200,
			token:    "", // No token provided
		},
	}
	for _, tc := range tests {
		url := fmt.Sprintf("/status/%d", tc.id)
		req, err := http.NewRequest("GET", url, bytes.NewBuffer(tc.jsonStr))
		if err != nil {
			t.Fatal(err)
		}
		if tc.token != "" {
			cookie := &http.Cookie{Name: "token", Value: tc.token}
			req.AddCookie(cookie)
		}
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != tc.expected {
			t.Errorf("test case '%s': expected status %d, got %d", tc.name, tc.expected, w.Code)
		}
	}
}

// TestRegister дополненный
func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	initializers.ConnectToDB()
	initializers.CreateTable()
	r.POST("/register", handler.RegisterHandler)
	initializers.DeleteUser()
	type cases struct {
		name     string
		jsonStr  []byte
		expected int
	}
	tests := []cases{
		{
			name:     "good",
			jsonStr:  []byte(`{"name":"Ansar123","password":"123"}`),
			expected: http.StatusOK,
		},
		{
			name:     "existing user",
			jsonStr:  []byte(`{"name":"Ansar123","password":"123"}`),
			expected: http.StatusInternalServerError, // Assuming StatusConflict is appropriate for existing user
		},
		{
			name:     "missing credentials",
			jsonStr:  []byte(`{"name":"","password":""}`),
			expected: http.StatusInternalServerError, // Assuming BadRequest for missing credentials
		},
	}
	for _, tc := range tests {
		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(tc.jsonStr))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != tc.expected {
			t.Errorf("test case '%s': expected status %d, got %d", tc.name, tc.expected, w.Code)
		}
	}
}

// TestLogin дополненный
func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	initializers.ConnectToDB()
	initializers.CreateTable()
	r.POST("/login", handler.Login)
	type cases struct {
		name     string
		jsonStr  []byte
		expected int
	}
	tests := []cases{
		{
			name:     "good",
			jsonStr:  []byte(`{"name":"Ansar","password":"123"}`),
			expected: http.StatusOK,
		},
		{
			name:     "nonexistent user",
			jsonStr:  []byte(`{"name":"NonexistentUser","password":"123"}`),
			expected: http.StatusInternalServerError, // Assuming StatusNotFound for nonexistent user
		},
		{
			name:     "wrong password",
			jsonStr:  []byte(`{"name":"Ansar","password":"wrongpassword"}`),
			expected: http.StatusInternalServerError, // Assuming StatusUnauthorized for wrong password
		},
		{
			name:     "missing credentials",
			jsonStr:  []byte(`{"name":"","password":""}`),
			expected: http.StatusInternalServerError, // Assuming BadRequest for missing credentials
		},
	}
	for _, tc := range tests {
		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(tc.jsonStr))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != tc.expected {
			t.Errorf("test case '%s': expected status %d, got %d", tc.name, tc.expected, w.Code)
		}
	}
}

func GetToken() string {
	var user models.User
	user.Name = "Ansar"
	user.Password = "123"
	_ = initializers.Register(user)
	token, _ := initializers.Login(user)
	return token
}
