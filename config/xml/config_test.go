package xml

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/ribeirohugo/go_config/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const configContent = `<config>
    <environment>dev</environment>
    <service>safesystem</service>
    <server>
        <host>localhost</host>
        <port>8080</port>
        <allowed_origins>http://localhost:8080</allowed_origins>
    </server>
    <token>
        <secret>token</secret>
        <max_age>100</max_age>
    </token>
    <mongodb>
        <database>database</database>
        <host>localhost</host>
        <password>password</password>
        <port>8080</port>
        <user>username</user>
    </mongodb>
    <mysql>
        <database>database</database>
        <host>localhost</host>
        <password>password</password>
        <port>8080</port>
        <user>username</user>
    </mysql>
    <postgres>
        <database>database</database>
        <host>localhost</host>
        <password>password</password>
        <port>8080</port>
        <user>username</user>
    </postgres>
    <tracer>
        <enabled>true</enabled>
        <jaeger_host>http://tracer.domain</jaeger_host>
    </tracer>
</config>
`

const configContentInvalid = `token = 123
[server]
host = localhost
port = "9399"
`

const configContentEmpty = `<config>
</config>
`

func TestLoad(t *testing.T) {
	const (
		environment  = "dev"
		service      = "safesystem"
		serverHost   = "localhost"
		serverPort   = 8080
		database     = "database"
		password     = "password"
		username     = "username"
		xmlLocalName = "config"
	)
	configTest := config.Config{
		XMLName: xml.Name{
			Local: xmlLocalName,
		},
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
		Tracer: config.Tracer{
			Enabled:    true,
			JaegerHost: "http://tracer.domain",
		},
		Environment: environment,
		Service:     service,
	}

	t.Run("should return a valid config from XML file", func(t *testing.T) {
		t.Run("with all fields", func(t *testing.T) {
			tempFile := createTempFile(t, configContent)

			cfg, err := Load(tempFile.Name())
			require.NoError(t, err)
			assert.Equal(t, configTest, cfg)

			closeFile(t, tempFile)
		})

		t.Run("without optional fields", func(t *testing.T) {
			expectedConfig := config.Config{
				XMLName: xml.Name{
					Local: xmlLocalName,
				},
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

			tempFile := createTempFile(t, configContentEmpty)

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

func createTempFile(t *testing.T, fileContent string) *os.File {
	t.Helper()

	tempFile, err := os.CreateTemp("", "test.toml")
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