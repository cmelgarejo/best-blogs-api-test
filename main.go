package main

import (
	"log"

	"go-rest-blog/bootstrap"
)

func main() {
	defaultPort := 8080
	if err := bootstrap.Init(defaultPort); err != nil {
		log.Fatalf("Service will be shutdown because error ocurred:  %+v", err.Error())
	}
}
