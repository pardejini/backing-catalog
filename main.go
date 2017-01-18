package main

import (
	"os"

	service "github.com/pardejini/backing-catalog/service"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	// Ordinarily we'd use a CF environment here, but we don't need it for
	// the fake data we're returning.
	server := service.NewServer()
	server.Run(":" + port)
}
