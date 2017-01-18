package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	service "github.com/pardejini/backing-catalog/service"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	appEnv, err := cfenv.Current()
	if err != nil {
		fmt.Println("CF Environment not detected.")
	}
	// Ordinarily we'd use a CF environment here, but we don't need it for
	// the fake data we're returning.
	// server := service.NewServer()

	server := service.NewServerFromCFEnv(appEnv)
	server.Run(":" + port)
}
