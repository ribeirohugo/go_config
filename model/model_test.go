package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	serverHost = "localhost"
	serverPort = 8080
	database   = "database"
	password   = "password"
	username   = "username"
)

func TestConfig_MongodbAddress(t *testing.T) {
	const (
		expectedAddress        = "mongodb://username:password@localhost:8080/database?authSource=admin&ssl=false"
		defaultMigrationsMongo = "file://migrations/mongo"
	)

	cfg := Config{
		MongoDb: Database{
			Host:           serverHost,
			Port:           serverPort,
			User:           username,
			Password:       password,
			Db:             database,
			MigrationsPath: defaultMigrationsMongo,
		},
	}

	t.Run("should return a valid MongoDB address", func(t *testing.T) {
		address := cfg.MongodbAddress()
		assert.Equal(t, expectedAddress, address)
	})
}

func TestConfig_MysqlAddress(t *testing.T) {
	const (
		expectedAddress        = "username:password@tcp(localhost:8080)/database"
		defaultMigrationsMysql = "file://migrations/mysql"
	)

	cfg := Config{
		MySql: Database{
			Host:           serverHost,
			Port:           serverPort,
			User:           username,
			Password:       password,
			Db:             database,
			MigrationsPath: defaultMigrationsMysql,
		},
	}

	t.Run("should return a valid MySQL address", func(t *testing.T) {
		address := cfg.MysqlAddress()
		assert.Equal(t, expectedAddress, address)
	})
}

func TestConfig_PostgresAddress(t *testing.T) {
	const (
		expectedAddress           = "postgres://username:password@localhost:8080/database?sslmode=disable"
		defaultMigrationsPostgres = "file://migrations/postgres"
	)

	cfg := Config{
		Postgres: Database{
			Host:           serverHost,
			Port:           serverPort,
			User:           username,
			Password:       password,
			Db:             database,
			MigrationsPath: defaultMigrationsPostgres,
		},
	}

	t.Run("should return a valid PostgreSQL address", func(t *testing.T) {
		address := cfg.PostgresAddress()
		assert.Equal(t, expectedAddress, address)
	})
}

func TestServer_GetAddress(t *testing.T) {
	const expectedAddress = "localhost:8080"

	server := Server{
		Host:           serverHost,
		Port:           serverPort,
		AllowedOrigins: []string{"http://localhost:8080"},
	}

	t.Run("should return a valid server address", func(t *testing.T) {
		address := server.GetAddress()
		assert.Equal(t, expectedAddress, address)
	})
}
