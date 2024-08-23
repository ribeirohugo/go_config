package env

import (
	"os"
	"testing"

	"github.com/ribeirohugo/go_config/config"

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
		environment = "dev"
		service     = "safesystem"
		serverHost  = "localhost"
		serverPort  = 8080
		database    = "database"
		password    = "password"
		username    = "username"
		auditHost   = "audit.domain"
		tracerHost  = "https://tracer.domain"
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
		Audit: config.Audit{
			Enabled: true,
			Host:    auditHost,
		},
		Tracer: config.Tracer{
			Enabled:    true,
			JaegerHost: tracerHost,
			Host:       tracerHost,
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
		err = os.Setenv("TRACER_ENABLED", "TRUE")
		require.NoError(t, err)
		err = os.Setenv("TRACER_HOST", tracerHost)
		require.NoError(t, err)
		err = os.Setenv("TRACER_JAEGER_HOST", tracerHost)
		require.NoError(t, err)
		err = os.Setenv("AUDIT_ENABLED", "TRUE")
		require.NoError(t, err)
		err = os.Setenv("AUDIT_HOST", auditHost)
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
				"TRACER_ENABLED",
				"TRACER_HOST",
				"TRACER_JAEGER_HOST",
				"AUDIT_ENABLED",
				"AUDIT_HOST",
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
				MigrationsPath: config.DefaultMigrationsMysql,
			},
			MongoDb: config.Database{
				MigrationsPath: config.DefaultMigrationsMongo,
			},
			Postgres: config.Database{
				MigrationsPath: config.DefaultMigrationsPostgres,
			},
			Tracer: config.Tracer{
				JaegerHost: config.DefaultJaegerHost,
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
			err := os.Setenv("TRACER_ENABLED", "error")
			require.NoError(t, err)
			defer func() {
				err = os.Unsetenv("TRACER_ENABLED")
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
