package main

import (
	"github.com/Rakanixu/elasticsearch/api/handler"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.elasticsearch"),
		micro.Version("latest"),
	)

	// Register Handler
	service.Server().Handle(
		service.Server().NewHandler(new(handler.Elasticsearch)),
	)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
