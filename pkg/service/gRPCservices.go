package service

import (
	"context"
	desc "github.com/RakhimovAns/FinalYandexTask/api/gen/api/service"
	"github.com/RakhimovAns/FinalYandexTask/initializers"
	"github.com/RakhimovAns/FinalYandexTask/models"
	"log"
	"strconv"
)

type Server struct {
	desc.UnimplementedCalculusServer
}

func (s *Server) Calculate(ctx context.Context, req *desc.Expression) (*desc.ID, error) {
	expression := models.Expression{
		Expression:   req.Expression,
		AddTime:      req.AddTime,
		DivideTime:   req.DivideTime,
		MultiplyTime: req.MultiplyTime,
		SubTime:      req.SubTime,
	}
	if err := IsValidate(expression); err != nil {
		log.Fatalf("IsValidate err: %v", err)
		return nil, err
	}

	id := initializers.CreateModel(expression)
	ID := &desc.ID{
		Id: id,
	}
	return ID, nil
}

func (s *Server) Status(ctx context.Context, req *desc.ID) (*desc.Result, error) {
	ID := req.Id
	expression := initializers.GetByID(ID)
	if expression.IsCounted {
		result := strconv.Itoa(int(expression.Result))
		Result := &desc.Result{
			Result: result,
		}
		return Result, nil
	}
	go func() {
		result, err := CountExpression(expression)
		if err != nil {
			log.Printf("Error counting expression: %v", err)
			return
		}
		initializers.SetResult(ID, result)
	}()
	Result := &desc.Result{
		Result: "is counting",
	}
	return Result, nil
}

func (s *Server) Register(ctx context.Context, req *desc.User) (*desc.Result, error) {
	user := models.User{
		Name:     req.Name,
		Password: req.Password,
	}
	err := initializers.Register(user)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	Result := &desc.Result{
		Result: "registered successfully, try to log in",
	}
	return Result, nil
}

func (s *Server) Login(ctx context.Context, req *desc.User) (*desc.Token, error) {
	user := models.User{
		Name:     req.Name,
		Password: req.Password,
	}
	token, err := initializers.Login(user)
	if err == models.ErrUserNotExist {
		return nil, err
	} else if err == models.ErrInvalidPassword {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return &desc.Token{Token: token}, nil
}

func (s *Server) Logout(ctx context.Context, req *desc.Empty) (*desc.Empty, error) {
	return &desc.Empty{}, nil
}
