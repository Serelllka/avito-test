package main

import (
	avito_test "avito-test"
	"avito-test/pkg/handler"
	"log"
	"os"
)

func main() {
	server := new(avito_test.Server)
	handle := handler.NewHandler()
	if err := server.Run(os.Args[1], handle.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
	println("Hello world!")
}
