package env

import (
	"os"
	"testing"

	"github.com/ribeirohugo/go_config/v2/pkg/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetNumber(t *testing.T) {
	t.Run("valid integer value", func(t *testing.T) {
		err := os.Setenv("TEST_KEY", "42")
		require.NoError(t, err)
		defer func() {
			err = os.Unsetenv("TEST_KEY")
			require.NoError(t, err)
		}()

		result, err := getNumber("TEST_KEY", defaultInt)
		assert.NoError(t, err)
		assert.Equal(t, 42, result)
	})

	t.Run("invalid integer value", func(t *testing.T) {
		err := os.Setenv("TEST_KEY", "not-a-number")
		require.NoError(t, err)
		defer func() {
			err = os.Unsetenv("TEST_KEY")
		}()

		result, err := getNumber("TEST_KEY", defaultInt)

		assert.Error(t, err)
		assert.Equal(t, defaultInt, result)
		assert.Contains(t, err.Error(), "invalid TEST_KEY")
	})

	t.Run("empty environment variable", func(t *testing.T) {
		err := os.Unsetenv("TEST_KEY")
		require.NoError(t, err)

		result, err := getNumber("TEST_KEY", defaultInt)
		require.NoError(t, err)
		assert.Equal(t, defaultInt, result)
	})
}

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
	)
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
		Environment: environment,
		Service:     service,
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
		err = os.Setenv("SERVICE", service)
		require.NoError(t, err)
		err = os.Setenv("ENVIRONMENT", environment)
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
				"SERVICE",
				"ENVIRONMENT",
			)
		}()

		cfg, err := Load()
		require.NoError(t, err)
		assert.Equal(t, expectedCfg, cfg)
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
			Tempo: config.ExternalService{
				Host: config.DefaultTempoHost,
			},
			Jaeger: config.ExternalService{
				Host: config.DefaultJaegerHost,
			},
			Token: config.Token{
				MaxAge: config.DefaultSessionMaxAge,
			},
		}

		cfg, err := Load()
		require.NoError(t, err)
		assert.Equal(t, expectedConfig, cfg)
	})

	t.Run("returns an error", func(t *testing.T) {
		t.Run("due to invalid int value", func(t *testing.T) {
			err := os.Setenv("SERVER_PORT", "error")
			require.NoError(t, err)
			defer func() {
				err = os.Unsetenv("SERVER_PORT")
				require.NoError(t, err)
			}()
			_, err = Load()
			assert.Error(t, err)
		})
		t.Run("due to invalid bool value", func(t *testing.T) {
			err := os.Setenv("JAEGER_ENABLED", "error")
			require.NoError(t, err)
			defer func() {
				err = os.Unsetenv("JAEGER_ENABLED")
				require.NoError(t, err)
			}()
			_, err = Load()
			assert.Error(t, err)
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
