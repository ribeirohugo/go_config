package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const configYmlContent = `environment: "dev"
service: "safesystem"

server:
  host: "localhost"
  port: 8080
  allowed_origins: ["http://localhost:8080"]

token:
  secret: "token"
  max_age: 100

mongodb:
  database: "database"
  host: "localhost"
  password: "password"
  port: 8080
  user: "username"

mysql:
  database: "database"
  host: "localhost"
  password: "password"
  port: 8080
  user: "username"

postgres:
  database: "database"
  host: "localhost"
  password: "password"
  port: 8080
  user: "username"

tracer:
  enabled: true
  jaeger_host: "http://tracer.domain"
`

const configYmlContentInvalid = `token: 123
server:
	host: localhost
	port: "9399"
`

func TestLoadYaml(t *testing.T) {
	t.Run("should return a valid config", func(t *testing.T) {
		t.Run("with all fields", func(t *testing.T) {
			tempFile := createTempFile(t, configYmlContent)

			cfg, err := LoadYaml(tempFile.Name())
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

			cfg, err := LoadYaml(tempFile.Name())
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
			tempFile := createTempFile(t, configYmlContentInvalid)

			cfg, err := LoadYaml(tempFile.Name())
			assert.Equal(t, Config{}, cfg)
			assert.Error(t, err)

			closeFile(t, tempFile)
		})
	})
}
