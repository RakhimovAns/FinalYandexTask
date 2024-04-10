package server

import (
	"fmt"
	desc "github.com/RakhimovAns/FinalYandexTask/api/gen/api/service"
	"github.com/RakhimovAns/FinalYandexTask/initializers"
	"github.com/RakhimovAns/FinalYandexTask/pkg/handlers"
	"github.com/RakhimovAns/FinalYandexTask/pkg/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"path/filepath"
)

const (
	grpcPort = 50051
)

func StartGRPCServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	desc.RegisterCalculusServer(s, &service.Server{})
	reflection.Register(s)
	log.Printf("gRPC server listening on port %d", grpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func StartGinServer() {
	log.Print("Start GIN SERVER")
	r := gin.Default()
	// Обработчик статических файлов
	r.Static("/static", "./static")

	// Обработчик для главной страницы
	r.GET("/", func(c *gin.Context) {
		c.File(filepath.Join("static", "index.html"))
	})
	r.GET("/reg", func(c *gin.Context) {
		c.File(filepath.Join("static", "register.html"))
	})
	r.GET("/log", func(c *gin.Context) {
		c.File(filepath.Join("static", "login.html"))
	})
	authorizedGroup := r.Group("/")
	authorizedGroup.Use(initializers.Authorized())
	{
		// Обработчик POST-запроса на вычисление выражения
		authorizedGroup.POST("/expression", handlers.PostExpression)

		// Обработчик для получения статуса выражения
		authorizedGroup.GET("/status/:id", handlers.GetStatus)
	}

	//Обработчик для регистрации
	r.POST("/register", handlers.RegisterHandler)
	r.POST("/login", handlers.Login)
	r.POST("/logout", handlers.Logout)
	r.Run()
}
