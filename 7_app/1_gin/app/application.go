package app

import (
	"learning/app/controllers"
)

type Application struct {
	Server *controllers.Server
}

func New() *Application {
	server := controllers.New()

	return &Application{
		Server: server,
	}
}

func (app *Application) Run() {
	// Initialize the application components here
	// For example, set up the database connection, start the server, etc.
	app.Server.Run()
}
