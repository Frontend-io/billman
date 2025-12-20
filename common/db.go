package common

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/rs/zerolog/log"
)

func ConnectDB() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Error().Err(err).Msg("Error loading .env file")
		return nil, err
	}

	// Retrieve database connection details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")

	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)

	log.Debug().Msg(connectionString)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Error().Err(err).Msg("Error opening db connection")
		return nil, err
	}

	// Verify the connection is active
	err = db.Ping()
	if err != nil {
		log.Error().Err(err).Msg("Unable to start db server")
		db.Close()
		return nil, err
	}

	log.Debug().Msg("Successfully connected to the DB!")

	return db, nil
}
