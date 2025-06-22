package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file: " + err.Error())
	}

	// Access config values
	fmt.Println("App Name:", viper.GetString("app_name"))
	fmt.Println("Port:", viper.GetInt("port"))
	fmt.Println("Debug Mode:", viper.GetBool("debug"))
	fmt.Println("Database Host:", viper.GetString("database.host"))

	viper.SetEnvPrefix("myapp") // Prefix for environment variables
	viper.AutomaticEnv()        // Automatically read environment variables

	// Override config values with environment variables
	// export MYAPP_PORT=9090
	// export MYAPP_DEBUG=true
	fmt.Println("Overridden Port: ", viper.GetInt("port"))
	fmt.Println("Overridden Database Mode: ", viper.GetString("debug"))
}
