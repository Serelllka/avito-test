package main

import (
	avito_test "avito-test"
	"avito-test/pkg/handler"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error while reading config: %s", err.Error())
	}
	server := new(avito_test.Server)
	handle := handler.NewHandler()
	if err := server.Run(viper.GetString("port"), handle.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
	println("Hello world!")
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
