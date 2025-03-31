package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ribeirohugo/go_config/v2/pkg/config"
)

const (
	defaultInt = 0
)

// Load loads configurations from a given json file path.
func Load() (config.Config, error) {
	serverPort, err := getNumber("SERVER_PORT", defaultInt)
	if err != nil {
		return config.Config{}, err
	}
	rawServerAllowedOrigins := os.Getenv("SERVER_ALLOWED_ORIGINS")
	var serverAllowedOrigins []string
	if rawServerAllowedOrigins != "" {
		serverAllowedOrigins = strings.Split(rawServerAllowedOrigins, ",")
	}
	tokenAge, err := getNumber("TOKEN_MAX_AGE", config.DefaultSessionMaxAge)
	if err != nil {
		return config.Config{}, err
	}
	auditEnabled, err := getBool("AUDIT_ENABLED")
	if err != nil {
		return config.Config{}, err
	}
	lokiEnabled, err := getBool("LOKI_ENABLED")
	if err != nil {
		return config.Config{}, err
	}
	lokiHost := os.Getenv("LOKI_HOST")
	if lokiHost == "" {
		lokiHost = config.DefaultLokiHost
	}
	prometheusEnabled, err := getBool("PROMETHEUS_ENABLED")
	if err != nil {
		return config.Config{}, err
	}
	tempoEnabled, err := getBool("TEMPO_ENABLED")
	if err != nil {
		return config.Config{}, err
	}
	tempoHost := os.Getenv("TEMPO_HOST")
	if tempoHost == "" {
		tempoHost = config.DefaultTempoHost
	}
	jaegerEnabled, err := getBool("JAEGER_ENABLED")
	if err != nil {
		return config.Config{}, err
	}
	jaegerHost := os.Getenv("JAEGER_HOST")
	if jaegerHost == "" {
		jaegerHost = config.DefaultJaegerHost
	}

	// Load Defaults
	mongoDBMigrationsPath := os.Getenv("MONGODB_MIGRATIONS_PATH")
	if mongoDBMigrationsPath == "" {
		mongoDBMigrationsPath = config.DefaultMigrationsMongo
	}
	mongoDBPort, err := getNumber("MONGODB_PORT", config.DefaultMongoPort)
	if err != nil {
		return config.Config{}, err
	}
	mySQLMigrationsPath := os.Getenv("MYSQL_MIGRATIONS_PATH")
	if mySQLMigrationsPath == "" {
		mySQLMigrationsPath = config.DefaultMigrationsMysql
	}
	mySQLPort, err := getNumber("MYSQL_PORT", config.DefaultMySQLPort)
	if err != nil {
		return config.Config{}, err
	}
	postgresMigrationsPath := os.Getenv("POSTGRES_MIGRATIONS_PATH")
	if postgresMigrationsPath == "" {
		postgresMigrationsPath = config.DefaultMigrationsPostgres
	}
	postgresPort, err := getNumber("POSTGRES_PORT", config.DefaultPostgresPort)
	if err != nil {
		return config.Config{}, err
	}

	// Load env variables
	cfg := config.Config{
		Server: config.Server{
			Host:           os.Getenv("SERVER_HOST"),
			Port:           serverPort,
			AllowedOrigins: serverAllowedOrigins,
		},
		Token: config.Token{
			MaxAge: tokenAge,
			Secret: os.Getenv("TOKEN_SECRET"),
		},
		MongoDb: config.Database{
			Host:           os.Getenv("MONGODB_HOST"),
			Port:           mongoDBPort,
			User:           os.Getenv("MONGODB_USER"),
			Password:       os.Getenv("MONGODB_PASSWORD"),
			Db:             os.Getenv("MONGODB_DATABASE"),
			MigrationsPath: mongoDBMigrationsPath,
		},
		MySql: config.Database{
			Host:           os.Getenv("MYSQL_HOST"),
			Port:           mySQLPort,
			User:           os.Getenv("MYSQL_USER"),
			Password:       os.Getenv("MYSQL_PASSWORD"),
			Db:             os.Getenv("MYSQL_DATABASE"),
			MigrationsPath: mySQLMigrationsPath,
		},
		Postgres: config.Database{
			Host:           os.Getenv("POSTGRES_HOST"),
			Port:           postgresPort,
			User:           os.Getenv("POSTGRES_USER"),
			Password:       os.Getenv("POSTGRES_PASSWORD"),
			Db:             os.Getenv("POSTGRES_DATABASE"),
			MigrationsPath: postgresMigrationsPath,
		},
		Audit: config.ExternalService{
			Enabled: auditEnabled,
			Host:    os.Getenv("AUDIT_HOST"),
			Token:   os.Getenv("AUDIT_TOKEN"),
		},
		Loki: config.ExternalService{
			Enabled: lokiEnabled,
			Host:    lokiHost,
			Token:   os.Getenv("LOKI_TOKEN"),
		},
		Prometheus: config.ExternalService{
			Enabled: prometheusEnabled,
			Host:    os.Getenv("PROMETHEUS_HOST"),
			Token:   os.Getenv("PROMETHEUS_TOKEN"),
		},
		Tempo: config.ExternalService{
			Enabled: tempoEnabled,
			Host:    tempoHost,
			Token:   os.Getenv("TEMPO_TOKEN"),
		},
		Jaeger: config.ExternalService{
			Enabled: jaegerEnabled,
			Host:    jaegerHost,
			Token:   os.Getenv("JAEGER_TOKEN"),
		},
		Environment: os.Getenv("ENVIRONMENT"),
		Service:     os.Getenv("SERVICE"),
	}

	return cfg, nil
}

func getNumber(key string, defaultVal int) (int, error) {
	rawIntValue := os.Getenv(key)
	if rawIntValue == "" {
		return defaultVal, nil
	}
	intValue, err := strconv.Atoi(rawIntValue)
	if err != nil {
		return defaultInt, fmt.Errorf("invalid %s int value: %s", key, err.Error())
	}
	return intValue, nil
}

func getBool(key string) (bool, error) {
	rawBoolValue := os.Getenv(key)
	switch rawBoolValue {
	case "1", "true", "TRUE", "True":
		return true, nil
	case "0", "false", "FALSE", "False", "":
		return false, nil
	}
	return false, fmt.Errorf("invalid %s bool value: %s", key, rawBoolValue)
}
