package main

import (
	"context"
	"fmt"
	desc "github.com/RakhimovAns/FinalYandexTask/api/gen/api/service"
	"github.com/RakhimovAns/FinalYandexTask/initializers"
	"github.com/RakhimovAns/FinalYandexTask/models"
	"github.com/RakhimovAns/FinalYandexTask/pkg/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"time"
)

const (
	grpcPort      = 50051
	httpPort      = 8080
	grpcServerURL = "localhost:50051"
)

type server struct {
	desc.UnimplementedCalculusServer
}

func (s *server) Calculate(ctx context.Context, req *desc.Expression) (*desc.ID, error) {
	expression := models.Expression{
		Expression:   req.Expression,
		AddTime:      req.AddTime,
		DivideTime:   req.DivideTime,
		MultiplyTime: req.MultiplyTime,
		SubTime:      req.SubTime,
	}
	if err := service.IsValidate(expression); err != nil {
		log.Fatalf("IsValidate err: %v", err)
		return nil, err
	}

	id := initializers.CreateModel(expression)
	ID := &desc.ID{
		Id: id,
	}
	return ID, nil
}

func (s *server) Status(ctx context.Context, req *desc.ID) (*desc.Result, error) {
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
		result, err := service.CountExpression(expression)
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

func (s *server) Register(ctx context.Context, req *desc.User) (*desc.Result, error) {
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

func Logout(c *gin.Context) {
	// Получаем список куки из запроса
	cookies := c.Request.Cookies()

	// Проходимся по всем кукам и устанавливаем MaxAge равным -1 для удаления
	for _, cookie := range cookies {
		newCookie := &http.Cookie{
			Name:   cookie.Name,
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		c.SetCookie(cookie.Name, "", -1, "/", "", false, true)
		c.SetCookie(newCookie.Name, newCookie.Value, newCookie.MaxAge, newCookie.Path, newCookie.Domain, newCookie.Secure, newCookie.HttpOnly)
	}

	// Выполняем перенаправление на страницу входа
	c.Redirect(http.StatusSeeOther, "/log")
}
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data: " + err.Error()})
		return
	}
	token, err := initializers.Login(user)
	if err == models.ErrUserNotExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No account with this name"})
		return
	} else if err == models.ErrInvalidPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords don't match"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	livingTime := 60 * time.Minute
	expiration := time.Now().Add(livingTime)
	// Set token in cookie

	cookie := http.Cookie{Name: "token", Value: url.QueryEscape(token), Expires: expiration}
	http.SetCookie(c.Writer, &cookie)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func main() {
	initializers.ConnectToDB()
	initializers.CreateTable()
	startGRPCServer()
}
func startGRPCServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	desc.RegisterCalculusServer(s, &server{})
	reflection.Register(s)
	log.Printf("gRPC server listening on port %d", grpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startGinServer() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.GET("/reg", func(c *gin.Context) {
		c.File(filepath.Join("static", "register.html"))
	})
	r.GET("/log", func(c *gin.Context) {
		c.File(filepath.Join("static", "login.html"))
	})
	//	r.POST("/register", Register)
	r.POST("/login", Login)
	r.POST("/logout", Logout)
	r.Run()
}
