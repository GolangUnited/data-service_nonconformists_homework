package main

import (
	"fmt"
	"golang-united-homework/config"
	"golang-united-homework/pkg/api"
	"golang-united-homework/pkg/database"
	"golang-united-homework/pkg/repositories"
	"golang-united-homework/pkg/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	err = database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()

	err = database.DB.AutoMigrate(&repositories.Homework{})
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	homeworkServer := &service.Homework{}

	api.RegisterHomeworkServer(grpcServer, homeworkServer)

	listener, err := net.Listen(config.PROTOCOL_TCP, fmt.Sprintf(":%s", config.PORT_8080))
	if err != nil {
		log.Fatal(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
