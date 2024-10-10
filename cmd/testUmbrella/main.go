package main

import (
	"log"
	"testUmbrella/internal/handler"
)

func main() {
	if err := handler.StartServer(); err != nil {
		log.Fatal(err)
	}
}
