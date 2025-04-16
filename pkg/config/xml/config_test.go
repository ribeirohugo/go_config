package xml

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ribeirohugo/go_config/v2/pkg/config"
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
    <audit>
        <enabled>true</enabled>
        <host>audit.domain</host>
        <token>audit.token</token>
    </audit>
    <loki>
        <enabled>true</enabled>
        <host>loki.domain</host>
        <token>loki.token</token>
    </loki>
    <prometheus>
        <enabled>true</enabled>
        <host>prometheus.domain</host>
        <token>prometheus.token</token>
    </prometheus>
    <tempo>
        <enabled>true</enabled>
        <host>tempo.domain</host>
        <token>tempo.token</token>
    </tempo>
    <jaeger>
        <enabled>true</enabled>
        <host>jaeger.domain</host>
        <token>jaeger.token</token>
    </jaeger>
    <redis>
        <enabled>true</enabled>
        <host>redis.domain</host>
        <token>redis.token</token>
    </redis>
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
		environment     = "dev"
		service         = "safesystem"
		serverHost      = "localhost"
		serverPort      = 8080
		database        = "database"
		password        = "password"
		username        = "username"
		xmlLocalName    = "config"
		auditHost       = "audit.domain"
		auditToken      = "audit.token"
		lokiHost        = "loki.domain"
		lokiToken       = "loki.token"
		prometheusHost  = "prometheus.domain"
		prometheusToken = "prometheus.token"
		tempoHost       = "tempo.domain"
		tempoToken      = "tempo.token"
		jaegerHost      = "jaeger.domain"
		jaegerToken     = "jaeger.token"
		redisHost       = "redis.domain"
		redisToken      = "redis.token"
	)
	configTest := config.XML{
		XMLName: xml.Name{
			Local: xmlLocalName,
		},
		Config: config.Config{
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
		},
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
			expectedConfig := config.XML{
				XMLName: xml.Name{
					Local: xmlLocalName,
				},
				Config: config.Config{
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
			assert.Equal(t, config.XML{}, cfg)
			assert.Error(t, err)
		})

		t.Run("invalid file content", func(t *testing.T) {
			tempFile := createTempFile(t, configContentInvalid)

			cfg, err := Load(tempFile.Name())
			assert.Equal(t, config.XML{}, cfg)
			assert.Error(t, err)

			closeFile(t, tempFile)
		})
	})
}

func TestLoadContent(t *testing.T) {
	const (
		environment     = "dev"
		service         = "safesystem"
		serverHost      = "localhost"
		serverPort      = 8080
		database        = "database"
		password        = "password"
		username        = "username"
		xmlLocalName    = "config"
		auditHost       = "audit.domain"
		auditToken      = "audit.token"
		lokiHost        = "loki.domain"
		lokiToken       = "loki.token"
		prometheusHost  = "prometheus.domain"
		prometheusToken = "prometheus.token"
		tempoHost       = "tempo.domain"
		tempoToken      = "tempo.token"
		jaegerHost      = "jaeger.domain"
		jaegerToken     = "jaeger.token"
		redisHost       = "redis.domain"
		redisToken      = "redis.token"
	)
	configTest := config.XML{
		XMLName: xml.Name{
			Local: xmlLocalName,
		},
		Config: config.Config{
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
		},
	}

	t.Run("should return a valid config from xml", func(t *testing.T) {
		t.Run("with all fields", func(t *testing.T) {
			cfg, err := LoadContent([]byte(configContent))
			require.NoError(t, err)
			assert.Equal(t, configTest, cfg)
		})

		t.Run("without optional fields", func(t *testing.T) {
			expectedConfig := config.XML{
				XMLName: xml.Name{
					Local: xmlLocalName,
				},
				Config: config.Config{
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
				},
			}

			cfg, err := LoadContent([]byte(configContentEmpty))
			require.NoError(t, err)
			assert.Equal(t, expectedConfig, cfg)
		})
	})

	t.Run("with error return", func(t *testing.T) {
		cfg, err := LoadContent([]byte(configContentInvalid))
		assert.Equal(t, config.XML{}, cfg)
		assert.Error(t, err)
	})
}

func createTempFile(t *testing.T, fileContent string) *os.File {
	t.Helper()

	tempFile, err := os.CreateTemp("", "test.xml")
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
