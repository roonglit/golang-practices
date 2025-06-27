package controllers_test

import (
	"context"
	"fmt"
	"learning/app/models"
	"learning/config"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/roonglit/credentials/pkg/credentials"
	"github.com/rs/zerolog/log"
)

var testConfig *config.Config
var testStore models.Store
var connPool *pgxpool.Pool

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Suite")
}

var _ = BeforeSuite(func() {
	gin.SetMode(gin.TestMode)

	testConfig = loadConfig()

	testStore = connectDb(testConfig)
})

var _ = AfterSuite(func() {
	// Expect(dbRunner.Stop()).To(Succeed())
})

var _ = AfterEach(func() {
	clearDatabase(connPool)
})

func loadConfig() *config.Config {
	// Initialize the ConfigReader with the correct config folder path
	reader := credentials.NewConfigReader("./../../config")

	// User-defined configuration struct
	var config config.Config

	// Read configuration with mode "debug" or "release"
	if err := reader.Read(gin.Mode(), &config); err != nil {
		log.Fatal().Err(err).Msgf("Failed to read configuration for mode: %s", gin.Mode())
	}

	// Debug: Print the loaded config
	// fmt.Printf("Loaded config - DBUri: %s, ServerAddress: %s\n", config.DBUri, config.ServerAddress)

	return &config
}

func connectDb(config *config.Config) models.Store {
	dbConfig, err := pgxpool.ParseConfig(config.DBUri)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to parse DB_URI")
	}

	connPool, err = pgxpool.New(context.Background(), dbConfig.ConnString())
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := models.NewStore(connPool)
	return store
}

func clearDatabase(pool *pgxpool.Pool) error {
	ctx := context.Background()

	tableNames := []string{
		"todos",
	}

	var truncateStmt string
	for _, tableName := range tableNames {
		truncateStmt += fmt.Sprintf("TRUNCATE %s CASCADE; ", tableName)
	}

	// Execute the truncate statement
	if _, err := pool.Exec(ctx, truncateStmt); err != nil {
		return fmt.Errorf("error executing truncate: %w", err)
	}

	return nil
}
