package app

import (
	"context"
	"fmt"
	"learning/app/controllers"
	"learning/app/models"
	"learning/app/tasks"
	"learning/config"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
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

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

	server := controllers.New(
		controllers.SetConfig(config),
		controllers.SetStore(store),
		controllers.SetAsynqClient(client),
	)

	go StartWorker()

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

func StartWorker() {
	fmt.Print("StartWorker called")
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeEmailSend, tasks.HandleEmailTask)

	if err := srv.Run(mux); err != nil {
		fmt.Printf("could not run server: %v", err)
	}
}
