package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const configContent = `environment = "dev"
service = "safesystem"

[server]
host = "localhost"
port = 8080
allowed_origins = ['http://localhost:8080']

[token]
secret = "token"
max_age = 100

[mongodb]
database = "database"
host = "localhost"
password = "password"
port = 8080
user = "username"

[mysql]
database = "database"
host = "localhost"
password = "password"
port = 8080
user = "username"

[postgres]
database = "database"
host = "localhost"
password = "password"
port = 8080
user = "username"

[tracer]
enabled = true
jaeger_host = "http://tracer.domain"
`

const configContentInvalid = `token = 123
[server]
host = localhost
port = "9399"
`

const (
	environment = "dev"
	service     = "safesystem"
	serverHost  = "localhost"
	serverPort  = 8080
	database    = "database"
	password    = "password"
	username    = "username"
)

var configTest = Config{
	Server: Server{
		Host:           serverHost,
		Port:           serverPort,
		AllowedOrigins: []string{"http://localhost:8080"},
	},
	Token: Token{
		MaxAge: 100,
		Secret: "token",
	},
	MongoDb: Database{
		Host:           serverHost,
		Port:           serverPort,
		User:           username,
		Password:       password,
		Db:             database,
		MigrationsPath: defaultMigrationsMongo,
	},
	MySql: Database{
		Host:           serverHost,
		Port:           serverPort,
		User:           username,
		Password:       password,
		Db:             database,
		MigrationsPath: defaultMigrationsMysql,
	},
	Postgres: Database{
		Host:           serverHost,
		Port:           serverPort,
		User:           username,
		Password:       password,
		Db:             database,
		MigrationsPath: defaultMigrationsPostgres,
	},
	Tracer: Tracer{
		Enabled:    true,
		JaegerHost: "http://tracer.domain",
	},
	Environment: environment,
	Service:     service,
}

func TestLoad(t *testing.T) {
	t.Run("should return a valid config", func(t *testing.T) {
		t.Run("with all fields", func(t *testing.T) {
			tempFile := createTempFile(t, configContent)

			cfg, err := Load(tempFile.Name())
			require.NoError(t, err)
			assert.Equal(t, configTest, cfg)

			closeFile(t, tempFile)
		})

		t.Run("without optional fields", func(t *testing.T) {
			expectedConfig := Config{
				MySql: Database{
					MigrationsPath: defaultMigrationsMysql,
				},
				MongoDb: Database{
					MigrationsPath: defaultMigrationsMongo,
				},
				Postgres: Database{
					MigrationsPath: defaultMigrationsPostgres,
				},
				Tracer: Tracer{
					JaegerHost: defaultJaegerHost,
				},
				Token: Token{
					MaxAge: defaultSessionMaxAge,
				},
			}

			tempFile := createTempFile(t, "")

			cfg, err := Load(tempFile.Name())
			require.NoError(t, err)
			assert.Equal(t, expectedConfig, cfg)

			closeFile(t, tempFile)
		})
	})

	t.Run("with error return", func(t *testing.T) {
		t.Run("file doesn't exist", func(t *testing.T) {
			cfg, err := Load("")
			assert.Equal(t, Config{}, cfg)
			assert.Error(t, err)
		})

		t.Run("invalid file content", func(t *testing.T) {
			tempFile := createTempFile(t, configContentInvalid)

			cfg, err := Load(tempFile.Name())
			assert.Equal(t, Config{}, cfg)
			assert.Error(t, err)

			closeFile(t, tempFile)
		})
	})
}

func TestConfig_MongodbAddress(t *testing.T) {
	const expectedAddress = "mongodb://username:password@localhost:8080/database?authSource=admin&ssl=false"

	t.Run("should return a valid MongoDB address", func(t *testing.T) {
		address := configTest.MongodbAddress()
		assert.Equal(t, expectedAddress, address)
	})
}

func TestConfig_MysqlAddress(t *testing.T) {
	const expectedAddress = "username:password@tcp(localhost:8080)/database"

	t.Run("should return a valid MySQL address", func(t *testing.T) {
		address := configTest.MysqlAddress()
		assert.Equal(t, expectedAddress, address)
	})
}

func TestConfig_PostgresAddress(t *testing.T) {
	const expectedAddress = "postgres://username:password@localhost:8080/database?sslmode=disable"

	t.Run("should return a valid PostgreSQL address", func(t *testing.T) {
		address := configTest.PostgresAddress()
		assert.Equal(t, expectedAddress, address)
	})
}

func TestServer_GetAddress(t *testing.T) {
	const expectedAddress = "localhost:8080"

	t.Run("should return a valid server address", func(t *testing.T) {
		address := configTest.Server.GetAddress()
		assert.Equal(t, expectedAddress, address)
	})
}

func createTempFile(t *testing.T, fileContent string) *os.File {
	t.Helper()

	tempFile, err := os.CreateTemp("", "config.toml")
	require.NoError(t, err)

	_, err = tempFile.WriteString(fileContent)
	require.NoError(t, err)

	return tempFile
}

func closeFile(t *testing.T, file *os.File) {
	t.Helper()

	err := file.Close()
	require.NoError(t, err)

	err = os.Remove(file.Name())
	require.NoError(t, err)
}
