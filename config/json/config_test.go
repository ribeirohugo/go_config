package json

import (
	"os"
	"testing"

	"github.com/ribeirohugo/go_config/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const configContent = `{
  "environment": "dev",
  "service": "safesystem",
  "server": {
    "host": "localhost",
    "port": 8080,
    "allowed_origins": ["http://localhost:8080"]
  },
  "token": {
    "secret": "token",
    "max_age": 100
  },
  "mongodb": {
    "database": "database",
    "host": "localhost",
    "password": "password",
    "port": 8080,
    "user": "username"
  },
  "mysql": {
    "database": "database",
    "host": "localhost",
    "password": "password",
    "port": 8080,
    "user": "username"
  },
  "postgres": {
    "database": "database",
    "host": "localhost",
    "password": "password",
    "port": 8080,
    "user": "username"
  },
  "audit": {
    "enabled": true,
    "host": "audit.domain"
  },
  "loki": {
    "enabled": true,
    "host": "loki.domain",
    "token": "loki.token"
  },
  "tracer": {
    "enabled": true,
    "jaeger_host": "https://tracer.domain",
    "host": "https://tracer.domain"
  }
}`

const configContentInvalid = `{
  "token": 123,
  "server": {
    "host": "localhost",
    "port": "9399"
  }
}`

func TestLoad(t *testing.T) {
	const (
		environment = "dev"
		service     = "safesystem"
		serverHost  = "localhost"
		serverPort  = 8080
		database    = "database"
		password    = "password"
		username    = "username"
		auditHost   = "audit.domain"
		lokiHost    = "loki.domain"
		lokiToken   = "loki.token"
		tracerHost  = "https://tracer.domain"
	)
	configTest := config.Config{
		Server: config.Server{
			Host:           serverHost,
			Port:           serverPort,
			AllowedOrigins: []string{"http://localhost:8080"},
		},
		Token: config.Token{
			MaxAge: 100,
			Secret: "token",
		},
		MongoDb: config.Database{
			Host:           serverHost,
			Port:           serverPort,
			User:           username,
			Password:       password,
			Db:             database,
			MigrationsPath: config.DefaultMigrationsMongo,
		},
		MySql: config.Database{
			Host:           serverHost,
			Port:           serverPort,
			User:           username,
			Password:       password,
			Db:             database,
			MigrationsPath: config.DefaultMigrationsMysql,
		},
		Postgres: config.Database{
			Host:           serverHost,
			Port:           serverPort,
			User:           username,
			Password:       password,
			Db:             database,
			MigrationsPath: config.DefaultMigrationsPostgres,
		},
		Audit: config.Audit{
			Enabled: true,
			Host:    auditHost,
		},
		Loki: config.ExternalService{
			Enabled: true,
			Host:    lokiHost,
			Token:   lokiToken,
		},
		Tracer: config.Tracer{
			Enabled:    true,
			JaegerHost: tracerHost,
			Host:       tracerHost,
		},
		Environment: environment,
		Service:     service,
	}

	t.Run("should return a valid config from json file", func(t *testing.T) {
		t.Run("with all fields", func(t *testing.T) {
			tempFile := createTempFile(t, configContent)

			cfg, err := Load(tempFile.Name())
			require.NoError(t, err)
			assert.Equal(t, configTest, cfg)

			closeFile(t, tempFile)
		})

		t.Run("without optional fields", func(t *testing.T) {
			expectedConfig := config.Config{
				MySql: config.Database{
					Port:           config.DefaultMySqlPort,
					MigrationsPath: config.DefaultMigrationsMysql,
				},
				MongoDb: config.Database{
					Port:           config.DefaultMongoPort,
					MigrationsPath: config.DefaultMigrationsMongo,
				},
				Postgres: config.Database{
					Port:           config.DefaultPostgresPort,
					MigrationsPath: config.DefaultMigrationsPostgres,
				},
				Audit: config.Audit{
					Enabled: false,
				},
				Loki: config.ExternalService{
					Enabled: false,
					Host:    config.DefaultLokiHost,
				},
				Tracer: config.Tracer{
					Enabled:    false,
					JaegerHost: config.DefaultJaegerHost,
				},
				Token: config.Token{
					MaxAge: config.DefaultSessionMaxAge,
				},
			}

			tempFile := createTempFile(t, "{}")

			cfg, err := Load(tempFile.Name())
			require.NoError(t, err)
			assert.Equal(t, expectedConfig, cfg)

			closeFile(t, tempFile)
		})
	})

	t.Run("with error return", func(t *testing.T) {
		t.Run("file doesn't exist", func(t *testing.T) {
			cfg, err := Load("")
			assert.Equal(t, config.Config{}, cfg)
			assert.Error(t, err)
		})

		t.Run("invalid file content", func(t *testing.T) {
			tempFile := createTempFile(t, configContentInvalid)

			cfg, err := Load(tempFile.Name())
			assert.Equal(t, config.Config{}, cfg)
			assert.Error(t, err)

			closeFile(t, tempFile)
		})
	})
}

func TestLoadContent(t *testing.T) {
	const (
		environment = "dev"
		service     = "safesystem"
		serverHost  = "localhost"
		serverPort  = 8080
		database    = "database"
		password    = "password"
		username    = "username"
		auditHost   = "audit.domain"
		lokiHost    = "loki.domain"
		lokiToken   = "loki.token"
		tracerHost  = "https://tracer.domain"
	)
	configTest := config.Config{
		Server: config.Server{
			Host:           serverHost,
			Port:           serverPort,
			AllowedOrigins: []string{"http://localhost:8080"},
		},
		Token: config.Token{
			MaxAge: 100,
			Secret: "token",
		},
		MongoDb: config.Database{
			Host:           serverHost,
			Port:           serverPort,
			User:           username,
			Password:       password,
			Db:             database,
			MigrationsPath: config.DefaultMigrationsMongo,
		},
		MySql: config.Database{
			Host:           serverHost,
			Port:           serverPort,
			User:           username,
			Password:       password,
			Db:             database,
			MigrationsPath: config.DefaultMigrationsMysql,
		},
		Postgres: config.Database{
			Host:           serverHost,
			Port:           serverPort,
			User:           username,
			Password:       password,
			Db:             database,
			MigrationsPath: config.DefaultMigrationsPostgres,
		},
		Audit: config.Audit{
			Enabled: true,
			Host:    auditHost,
		},
		Loki: config.ExternalService{
			Enabled: true,
			Host:    lokiHost,
			Token:   lokiToken,
		},
		Tracer: config.Tracer{
			Enabled:    true,
			JaegerHost: tracerHost,
			Host:       tracerHost,
		},
		Environment: environment,
		Service:     service,
	}

	t.Run("should return a valid config from json", func(t *testing.T) {
		t.Run("with all fields", func(t *testing.T) {
			cfg, err := LoadContent([]byte(configContent))
			require.NoError(t, err)
			assert.Equal(t, configTest, cfg)
		})

		t.Run("without optional fields", func(t *testing.T) {
			expectedConfig := config.Config{
				MySql: config.Database{
					Port:           config.DefaultMySqlPort,
					MigrationsPath: config.DefaultMigrationsMysql,
				},
				MongoDb: config.Database{
					Port:           config.DefaultMongoPort,
					MigrationsPath: config.DefaultMigrationsMongo,
				},
				Postgres: config.Database{
					Port:           config.DefaultPostgresPort,
					MigrationsPath: config.DefaultMigrationsPostgres,
				},
				Loki: config.ExternalService{
					Host: config.DefaultLokiHost,
				},
				Tracer: config.Tracer{
					JaegerHost: config.DefaultJaegerHost,
				},
				Token: config.Token{
					MaxAge: config.DefaultSessionMaxAge,
				},
			}

			cfg, err := LoadContent([]byte("{}"))
			require.NoError(t, err)
			assert.Equal(t, expectedConfig, cfg)
		})
	})

	t.Run("with error return", func(t *testing.T) {
		cfg, err := LoadContent([]byte(""))
		assert.Equal(t, config.Config{}, cfg)
		assert.Error(t, err)
	})
}

func createTempFile(t *testing.T, fileContent string) *os.File {
	t.Helper()

	tempFile, err := os.CreateTemp("", "test.json")
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
