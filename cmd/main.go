package main

import (
	"github.com/RakhimovAns/FinalYandexTask/initializers"
	"github.com/RakhimovAns/FinalYandexTask/pkg/server"
)

func main() {

}

func init() {
	initializers.ConnectToDB()
	initializers.CreateTable()
	go server.StartGRPCServer()
	server.StartGinServer()
}
