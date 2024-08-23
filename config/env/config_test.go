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
	t.Run("with valid environment variables", func(t *testing.T) {
		err := os.Setenv("SERVER_PORT", "8080")
		require.NoError(t, err)
		err = os.Setenv("SERVER_ALLOWED_ORIGINS", "http://example.com,http://another.com")
		require.NoError(t, err)
		err = os.Setenv("TOKEN_AGE", "3600")
		require.NoError(t, err)
		err = os.Setenv("MONGODB_PORT", "27017")
		require.NoError(t, err)
		err = os.Setenv("MYSQL_PORT", "3306")
		require.NoError(t, err)
		err = os.Setenv("POSTGRES_PORT", "5432")
		require.NoError(t, err)
		defer func() {
			unsetEnvVars(t,
				"SERVER_PORT",
				"SERVER_ALLOWED_ORIGINS",
				"TOKEN_AGE",
				"MONGODB_PORT",
				"MYSQL_PORT",
				"POSTGRES_PORT",
			)
		}()

		cfg, err := Load()
		require.NoError(t, err)

		assert.Equal(t, 8080, cfg.Server.Port)
		assert.Equal(t, []string{"http://example.com", "http://another.com"}, cfg.Server.AllowedOrigins)
		assert.Equal(t, 3600, cfg.Token.MaxAge)
		assert.Equal(t, 27017, cfg.MongoDb.Port)
		assert.Equal(t, 3306, cfg.MySql.Port)
		assert.Equal(t, 5432, cfg.Postgres.Port)
		assert.Equal(t, config.DefaultMigrationsMongo, cfg.MongoDb.MigrationsPath)
		assert.Equal(t, config.DefaultMigrationsMysql, cfg.MySql.MigrationsPath)
		assert.Equal(t, config.DefaultMigrationsPostgres, cfg.Postgres.MigrationsPath)
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

	t.Run("load with invalid SERVER_PORT", func(t *testing.T) {
		// Set the environment variable
		os.Setenv("SERVER_PORT", "invalid-port")

		defer os.Unsetenv("SERVER_PORT")

		// Call the function
		_, err := Load()

		// Assert the results
		assert.Error(t, err)
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
