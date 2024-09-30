package main

import (
	"log"
	"webinar-service/internal/api"
)

func main() {
	if err := api.Serve(); err != nil {
		log.Fatal(err)
	}
}
