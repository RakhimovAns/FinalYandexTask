package main

import (
	"context"
	"fmt"
	desc "github.com/RakhimovAns/FinalYandexTask/api/gen/api/service"
	"github.com/RakhimovAns/FinalYandexTask/initializers"
	"github.com/RakhimovAns/FinalYandexTask/models"
	"github.com/RakhimovAns/FinalYandexTask/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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
