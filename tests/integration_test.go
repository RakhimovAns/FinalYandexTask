package tests

import (
	"bytes"
	"github.com/RakhimovAns/FinalYandexTask/initializers"
	"github.com/RakhimovAns/FinalYandexTask/models"
	"github.com/RakhimovAns/FinalYandexTask/pkg/handlers"
	"github.com/RakhimovAns/FinalYandexTask/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Проверка для сохранения в бд
func TestPostExpression(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	initializers.ConnectToDB()
	initializers.CreateTable()
	go server.StartGRPCServer()
	r.POST("/expression", handlers.PostExpression)
	jsonStr := []byte(`{"expression":"2+2","addTime":1,"divideTime":1,"subTime":1,"multiplyTime":1}`)
	req, err := http.NewRequest("POST", "/expression", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	token := GetToken()
	cookie := &http.Cookie{Name: "token", Value: token}
	req.AddCookie(cookie)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

// Проверка для подсчета
func TestStatusExpression(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	initializers.ConnectToDB()
	initializers.CreateTable()
	go server.StartGRPCServer()
	r.GET("/status/1", handlers.GetStatus)
	jsonStr := []byte(`{"id": 1}`)
	req, err := http.NewRequest("GET", "/status/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	token := GetToken()
	cookie := &http.Cookie{Name: "token", Value: token}
	req.AddCookie(cookie)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

// Проверка регистрации
func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	initializers.ConnectToDB()
	initializers.CreateTable()
	go server.StartGRPCServer()
	r.POST("/register", handlers.RegisterHandler)
	jsonStr := []byte(`{"name":"Ansar1","password":"123"}`)
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

// Проверка Входа
func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	initializers.ConnectToDB()
	initializers.CreateTable()
	go server.StartGRPCServer()
	r.POST("/login", handlers.Login)
	jsonStr := []byte(`{"name":"Ansar1","password":"123"}`)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func GetToken() string {
	var user models.User
	user.Name = "Ansar"
	user.Password = "123"
	token, _ := initializers.Login(user)
	return token
}
