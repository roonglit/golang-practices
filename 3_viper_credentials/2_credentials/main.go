package main

import (
	"fmt"
	"learning/config"

	"github.com/roonglit/credentials/pkg/credentials"
)

func main() {
	reader := credentials.NewConfigReader()

	var config config.Config

	if err := reader.Read("debug", &config); err != nil {
		panic(err)
	}

	fmt.Println("AppName:", config.AppName)
	fmt.Println("Port:", config.Port)
	fmt.Println("DBUri:", config.DBUri)
	fmt.Println("Config loaded successfully")

	if err := reader.Read("release", &config); err != nil {
		panic(err)
	}
	fmt.Println("AppName:", config.AppName)
	fmt.Println("Port:", config.Port)
	fmt.Println("DBUri:", config.DBUri)
	fmt.Println("Config loaded successfully in release mode")

	// Override config values with environment variables
	// Call this in your terminal before running the program
	// export APP_NAME="App Using Credentials (deployed)"
	// export PORT=3456
	// export DB_URI="postgres://user:password@localhost:5432/dbname?"
}
