package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type server struct {
	echo *echo.Echo
}

func newServer() *server {
	s := &server{
		echo: echo.New(),
	}

	s.InitRoutes()

	return s
}

func StartServer() error {

	srv := newServer()

	if err := initConfig(); err != nil {
		log.Fatalf("cannot start server %s", err)
		return err
	}

	port := viper.GetString("port")

	fmt.Println("server started")

	return http.ListenAndServe(port, srv.echo)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
