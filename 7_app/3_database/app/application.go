package app

import (
	"context"
	"learning/app/controllers"
	"learning/app/models"
	"learning/config"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/roonglit/credentials/pkg/credentials"
)

type Application struct {
	Server *controllers.Server
}

func New() *Application {
	config := loadConfig()

	// database connection
	store := connectDb(config)

	server := controllers.New(
		config,
		store,
	)

	return &Application{
		Server: server,
	}
}

func (app *Application) Run() {
	// Initialize the application components here
	// For example, set up the database connection, start the server, etc.
	app.Server.Run()
}

func loadConfig() *config.Config {
	reader := credentials.NewConfigReader()

	var config config.Config

	if err := reader.Read(gin.Mode(), &config); err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	return &config
}

func connectDb(config *config.Config) models.Store {
	dbConfig, err := pgxpool.ParseConfig(config.DBUri)
	if err != nil {
		// log.Fatal().Err(err).Msg("Unable to parse DB_URI")
	}

	connPool, err := pgxpool.New(context.Background(), dbConfig.ConnString())
	if err != nil {
		// log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := models.NewStore(connPool)

	return store
}
