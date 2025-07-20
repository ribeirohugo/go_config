package dotenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ribeirohugo/go_config/v2/pkg/config"
)

const configContent = `ENVIRONMENT=dev
SERVICE=safesystem

SERVER_HOST=localhost
SERVER_PORT=8080
SERVER_ALLOWED_ORIGINS=http://localhost:8080

TOKEN_SECRET=token
MAX_AGE=100

MONGODB_DATABASE=database
MONGODB_HOST=localhost
MONGODB_PASSWORD=password
MONGODB_PORT=8080
MONGODB_USER=username

MYSQL_DATABASE=database
MYSQL_HOST=localhost
MYSQL_PASSWORD=password
MYSQL_PORT=8080
MYSQL_USER=username

POSTGRES_DATABASE=database
POSTGRES_HOST=localhost
POSTGRES_PASSWORD=password
POSTGRES_PORT=8080
POSTGRES_USER=username

AUDIT_ENABLED=true
AUDIT_HOST=audit.domain
AUDIT_TOKEN=audit.token

LOKI_ENABLED=true
LOKI_HOST=loki.domain
LOKI_TOKEN=loki.token

PROMETHEUS_ENABLED=true
PROMETHEUS_HOST=prometheus.domain
PROMETHEUS_TOKEN=prometheus.token

TEMPO_ENABLED=true
TEMPO_HOST=tempo.domain
TEMPO_TOKEN=tempo.token

JAEGER_ENABLED=true
JAEGER_HOST=jaeger.domain
JAEGER_TOKEN=jaeger.token

REDIS_ENABLED=true
REDIS_HOST=redis.domain
REDIS_TOKEN=redis.token

TRACER_ENABLED=true
TRACER_HOST=https://tracer.domain
TRACER_TOKEN=https://tracer.domain

SETTINGS=setting1=value1,setting2=value2,setting3=value3
`

const configContentInvalid = `token=123
[server]
host=localhost
port=9399
`

func TestLoad(t *testing.T) {
	const (
		environment     = "dev"
		service         = "safesystem"
		serverHost      = "localhost"
		serverPort      = 8080
		database        = "database"
		password        = "password"
		username        = "username"
		auditHost       = "audit.domain"
		auditToken      = "audit.token"
		lokiHost        = "loki.domain"
		lokiToken       = "loki.token"
		prometheusHost  = "prometheus.host"
		prometheusToken = "prometheus.token"
		tempoHost       = "tempo.domain"
		tempoToken      = "tempo.token"
		jaegerHost      = "jaeger.domain"
		jaegerToken     = "jaeger.token"
		redisHost       = "redis.domain"
		redisToken      = "redis.token"
		envSettings     = "setting1=value1,setting2=value2,setting3=value3"
	)
	settings := map[string]string{
		"setting1": "value1",
		"setting2": "value2",
		"setting3": "value3",
	}
	expectedCfg := config.Config{
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
		Audit: config.ExternalService{
			Enabled: true,
			Host:    auditHost,
			Token:   auditToken,
		},
		Loki: config.ExternalService{
			Enabled: true,
			Host:    lokiHost,
			Token:   lokiToken,
		},
		Prometheus: config.ExternalService{
			Enabled: true,
			Host:    prometheusHost,
			Token:   prometheusToken,
		},
		Tempo: config.ExternalService{
			Enabled: true,
			Host:    tempoHost,
			Token:   tempoToken,
		},
		Jaeger: config.ExternalService{
			Enabled: true,
			Host:    jaegerHost,
			Token:   jaegerToken,
		},
		Redis: config.ExternalService{
			Enabled: true,
			Host:    redisHost,
			Token:   redisToken,
		},
		Environment: environment,
		Service:     service,
		Settings:    settings,
	}

	t.Run("with valid environment variables", func(t *testing.T) {
		err := os.Setenv("SERVER_HOST", serverHost)
		require.NoError(t, err)
		err = os.Setenv("SERVER_PORT", "8080")
		require.NoError(t, err)
		err = os.Setenv("SERVER_ALLOWED_ORIGINS", "http://localhost:8080")
		require.NoError(t, err)
		err = os.Setenv("TOKEN_MAX_AGE", "100")
		require.NoError(t, err)
		err = os.Setenv("TOKEN_SECRET", "token")
		require.NoError(t, err)
		err = os.Setenv("MONGODB_HOST", serverHost)
		require.NoError(t, err)
		err = os.Setenv("MONGODB_PORT", "8080")
		require.NoError(t, err)
		err = os.Setenv("MONGODB_USER", username)
		require.NoError(t, err)
		err = os.Setenv("MONGODB_PASSWORD", password)
		require.NoError(t, err)
		err = os.Setenv("MONGODB_DATABASE", database)
		require.NoError(t, err)
		err = os.Setenv("MYSQL_HOST", serverHost)
		require.NoError(t, err)
		err = os.Setenv("MYSQL_PORT", "8080")
		require.NoError(t, err)
		err = os.Setenv("MYSQL_USER", username)
		require.NoError(t, err)
		err = os.Setenv("MYSQL_PASSWORD", password)
		require.NoError(t, err)
		err = os.Setenv("MYSQL_DATABASE", database)
		require.NoError(t, err)
		err = os.Setenv("POSTGRES_HOST", serverHost)
		require.NoError(t, err)
		err = os.Setenv("POSTGRES_PORT", "8080")
		require.NoError(t, err)
		err = os.Setenv("POSTGRES_USER", username)
		require.NoError(t, err)
		err = os.Setenv("POSTGRES_PASSWORD", password)
		require.NoError(t, err)
		err = os.Setenv("POSTGRES_DATABASE", database)
		require.NoError(t, err)
		err = os.Setenv("AUDIT_ENABLED", "TRUE")
		require.NoError(t, err)
		err = os.Setenv("AUDIT_HOST", auditHost)
		require.NoError(t, err)
		err = os.Setenv("AUDIT_TOKEN", auditToken)
		require.NoError(t, err)
		err = os.Setenv("LOKI_ENABLED", "TRUE")
		require.NoError(t, err)
		err = os.Setenv("LOKI_HOST", lokiHost)
		require.NoError(t, err)
		err = os.Setenv("LOKI_TOKEN", lokiToken)
		require.NoError(t, err)
		err = os.Setenv("PROMETHEUS_ENABLED", "TRUE")
		require.NoError(t, err)
		err = os.Setenv("PROMETHEUS_HOST", prometheusHost)
		require.NoError(t, err)
		err = os.Setenv("PROMETHEUS_TOKEN", prometheusToken)
		require.NoError(t, err)
		err = os.Setenv("TEMPO_ENABLED", "TRUE")
		require.NoError(t, err)
		err = os.Setenv("TEMPO_HOST", tempoHost)
		require.NoError(t, err)
		err = os.Setenv("TEMPO_TOKEN", tempoToken)
		require.NoError(t, err)
		err = os.Setenv("JAEGER_ENABLED", "TRUE")
		require.NoError(t, err)
		err = os.Setenv("JAEGER_HOST", jaegerHost)
		require.NoError(t, err)
		err = os.Setenv("JAEGER_TOKEN", jaegerToken)
		require.NoError(t, err)
		err = os.Setenv("REDIS_ENABLED", "TRUE")
		require.NoError(t, err)
		err = os.Setenv("REDIS_HOST", redisHost)
		require.NoError(t, err)
		err = os.Setenv("REDIS_TOKEN", redisToken)
		require.NoError(t, err)
		err = os.Setenv("SERVICE", service)
		require.NoError(t, err)
		err = os.Setenv("ENVIRONMENT", environment)
		require.NoError(t, err)
		err = os.Setenv("SETTINGS", envSettings)
		require.NoError(t, err)
		defer func() {
			unsetEnvVars(t,
				"SERVER_HOST",
				"SERVER_PORT",
				"SERVER_ALLOWED_ORIGINS",
				"TOKEN_MAX_AGE",
				"TOKEN_SECRET",
				"MONGODB_HOST",
				"MONGODB_PORT",
				"MONGODB_USER",
				"MONGODB_PASSWORD",
				"MONGODB_DATABASE",
				"MYSQL_HOST",
				"MYSQL_PORT",
				"MYSQL_USER",
				"MYSQL_PASSWORD",
				"MYSQL_DATABASE",
				"POSTGRES_HOST",
				"POSTGRES_PORT",
				"POSTGRES_USER",
				"POSTGRES_PASSWORD",
				"POSTGRES_DATABASE",
				"AUDIT_ENABLED",
				"AUDIT_HOST",
				"AUDIT_TOKEN",
				"LOKI_ENABLED",
				"LOKI_HOST",
				"LOKI_TOKEN",
				"PROMETHEUS_ENABLED",
				"PROMETHEUS_HOST",
				"PROMETHEUS_TOKEN",
				"TEMPO_ENABLED",
				"TEMPO_HOST",
				"TEMPO_TOKEN",
				"JAEGER_ENABLED",
				"JAEGER_HOST",
				"JAEGER_TOKEN",
				"REDIS_ENABLED",
				"REDIS_HOST",
				"REDIS_TOKEN",
				"SERVICE",
				"ENVIRONMENT",
				"SETTINGS",
			)
		}()

		tempFile := createTempFile(t, configContent)

		cfg, err := Load(tempFile.Name())
		require.NoError(t, err)
		assert.Equal(t, expectedCfg, cfg)

		closeFile(t, tempFile)
	})

	t.Run("without optional fields", func(t *testing.T) {
		expectedConfig := config.Config{
			MySql: config.Database{
				Port:           config.DefaultMySQLPort,
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
			Tempo: config.ExternalService{
				Host: config.DefaultTempoHost,
			},
			Jaeger: config.ExternalService{
				Host: config.DefaultJaegerHost,
			},
			Redis: config.ExternalService{
				Host: config.DefaultRedisHost,
			},
			Token: config.Token{
				MaxAge: config.DefaultSessionMaxAge,
			},
			Settings: map[string]string{},
		}
		defer func() {
			unsetEnvVars(t,
				"SERVER_HOST",
				"SERVER_PORT",
				"SERVER_ALLOWED_ORIGINS",
				"TOKEN_MAX_AGE",
				"TOKEN_SECRET",
				"MONGODB_HOST",
				"MONGODB_PORT",
				"MONGODB_USER",
				"MONGODB_PASSWORD",
				"MONGODB_DATABASE",
				"MYSQL_HOST",
				"MYSQL_PORT",
				"MYSQL_USER",
				"MYSQL_PASSWORD",
				"MYSQL_DATABASE",
				"POSTGRES_HOST",
				"POSTGRES_PORT",
				"POSTGRES_USER",
				"POSTGRES_PASSWORD",
				"POSTGRES_DATABASE",
				"AUDIT_ENABLED",
				"AUDIT_HOST",
				"AUDIT_TOKEN",
				"LOKI_ENABLED",
				"LOKI_HOST",
				"LOKI_TOKEN",
				"PROMETHEUS_ENABLED",
				"PROMETHEUS_HOST",
				"PROMETHEUS_TOKEN",
				"TEMPO_ENABLED",
				"TEMPO_HOST",
				"TEMPO_TOKEN",
				"JAEGER_ENABLED",
				"JAEGER_HOST",
				"JAEGER_TOKEN",
				"REDIS_ENABLED",
				"REDIS_HOST",
				"REDIS_TOKEN",
				"SERVICE",
				"ENVIRONMENT",
				"SETTINGS",
			)
		}()

		tempFile := createTempFile(t, "")

		cfg, err := Load(tempFile.Name())
		require.NoError(t, err)
		assert.Equal(t, expectedConfig, cfg)

		closeFile(t, tempFile)
	})

	t.Run("returns an error", func(t *testing.T) {
		t.Run("due to invalid int value", func(t *testing.T) {
			err := os.Setenv("SERVER_PORT", "error")
			require.NoError(t, err)
			defer func() {
				err = os.Unsetenv("SERVER_PORT")
				require.NoError(t, err)
			}()
			_, err = Load("")
			assert.Error(t, err)
		})

		t.Run("due to invalid bool value", func(t *testing.T) {
			err := os.Setenv("JAEGER_ENABLED", "error")
			require.NoError(t, err)
			defer func() {
				err = os.Unsetenv("JAEGER_ENABLED")
				require.NoError(t, err)
			}()
			_, err = Load("")
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

func unsetEnvVars(t *testing.T, vars ...string) {
	t.Helper()
	for _, key := range vars {
		err := os.Unsetenv(key)
		require.NoError(t, err, "Failed to unset environment variable: %s", key)
	}
}

func createTempFile(t *testing.T, fileContent string) *os.File {
	t.Helper()

	tempFile, err := os.CreateTemp("", ".env")
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
