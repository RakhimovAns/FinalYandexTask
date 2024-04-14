package service

import (
	"context"
	"errors"
	desc "github.com/RakhimovAns/FinalYandexTask/api/gen/api/service"
	"github.com/RakhimovAns/FinalYandexTask/initializers"
	"github.com/RakhimovAns/FinalYandexTask/models"
	"testing"
)

func NewServer() *Server {
	return &Server{}
}

func TestRegister(t *testing.T) {
	server := NewServer()
	initializers.ConnectToDB()
	type test struct {
		name     string
		User     *desc.User
		expected error
	}
	cases := []test{
		{
			name:     "All Right",
			User:     &desc.User{Name: "Chuk12", Password: "123"},
			expected: nil,
		},
	}
	for _, tests := range cases {
		_, err := server.Register(context.Background(), tests.User)
		if errors.Is(err, tests.expected) {
			t.Errorf("Expected %v,result :%v", tests.expected, err)
		}
	}

}

func TestLogin(t *testing.T) {
	server := NewServer()
	initializers.ConnectToDB()
	type test struct {
		name     string
		user     *desc.User
		expected error
	}
	cases := []test{
		{
			name:     "Wrong Password",
			user:     &desc.User{Name: "Ansar", Password: "12345"},
			expected: models.ErrInvalidPassword,
		},
		{
			name:     "No user",
			user:     &desc.User{Name: "Check", Password: "12345"},
			expected: models.ErrUserNotExist,
		},
		{
			name:     "All Right",
			user:     &desc.User{Name: "Chuk", Password: "123"},
			expected: nil,
		},
	}

	for _, tests := range cases {
		_, err := server.Login(context.Background(), tests.user)
		if !errors.Is(err, tests.expected) {
			t.Errorf("Expected %v,result :%v", tests.expected, err)
		}
	}
}

func TestCalculate(t *testing.T) {
	server := NewServer()
	initializers.ConnectToDB()
	type test struct {
		name       string
		expression *desc.Expression
		expected   error
	}
	cases := []test{
		{
			name:       "Wrong Expression",
			expression: &desc.Expression{Expression: "1+2+"},
			expected:   models.ErrInvalidExpression,
		},
		{
			name:       "Good Expression",
			expression: &desc.Expression{Expression: "1+2"},
			expected:   nil,
		},
	}
	for _, tests := range cases {
		_, err := server.Calculate(context.Background(), tests.expression)
		if !errors.Is(err, tests.expected) {
			t.Errorf("Expected :%v,result :%v", tests.expected, err)
		}
	}
}

func TestStatus(t *testing.T) {
	server := NewServer()
	initializers.ConnectToDB()
	type test struct {
		name     string
		id       *desc.ID
		expected error
	}
	cases := []test{
		{
			name:     "Not existing ID",
			id:       &desc.ID{Id: 100},
			expected: nil,
		},
		{
			name:     "OK",
			id:       &desc.ID{Id: 23},
			expected: nil,
		},
	}
	for _, tests := range cases {
		_, err := server.Status(context.Background(), tests.id)
		if !errors.Is(err, tests.expected) {
			t.Errorf("Expected :%v,result :%v", tests.expected, err)
		}
	}
}

func TestLogout(t *testing.T) {
	server := NewServer()
	initializers.ConnectToDB()
	type test struct {
		name     string
		empty    *desc.Empty
		expected error
	}
	cases := []test{
		{
			name:     "Nothing Special",
			empty:    &desc.Empty{},
			expected: nil,
		},
	}
	for _, tests := range cases {
		_, err := server.Logout(context.Background(), tests.empty)
		if !errors.Is(err, tests.expected) {
			t.Errorf("Expected :%v,result :%v", tests.expected, err)
		}
	}
}
