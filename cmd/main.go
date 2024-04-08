package main

import (
	desc "github.com/RakhimovAns/FinalYandexTask/api/gen/api/service"
)

const (
	grpcPort      = 50051
	httpPort      = 8080
	grpcServerURL = "localhost:50051"
)

type server struct {
	desc.UnimplementedCalculusServer
}

//
//func (s *server) Calculate(ctx context.Context, req *desc.Expression) (*desc.ID, error) {
//
//}
