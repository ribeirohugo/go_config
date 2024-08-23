package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ribeirohugo/go_config/config"
)

// Load loads configurations from a given json file path.
func Load() (config.Config, error) {
	// Load defaults
	rawServerPort := os.Getenv("SERVER_PORT")
	serverPort, err := strconv.Atoi(rawServerPort)
	if err != nil {
		return config.Config{}, fmt.Errorf("invalid SERVER_PORT: %s", err.Error())
	}
	rawServerAllowedOrigins := os.Getenv("SERVER_ALLOWED_ORIGINS")
	var serverAllowedOrigins []string
	if rawServerAllowedOrigins != "" {
		serverAllowedOrigins = strings.Split(rawServerAllowedOrigins, ",")
	}
	rawTokenAge := os.Getenv("TOKEN_AGE")
	tokenAge, err := strconv.Atoi(rawTokenAge)
	if err != nil {
		return config.Config{}, fmt.Errorf("invalid TOKEN_AGE: %s", err.Error())
	}

	// Load Defaults
	mongoDBMigrationsPath := os.Getenv("MONGODB_MIGRATIONS_PATH")
	if mongoDBMigrationsPath == "" {
		mongoDBMigrationsPath = config.DefaultMigrationsMongo
	}
	rawMongoDBPort := os.Getenv("MONGODB_PORT")
	mongoDBPort, err := strconv.Atoi(rawMongoDBPort)
	if err != nil {
		return config.Config{}, fmt.Errorf("invalid MONGODB_PORT: %s", err.Error())
	}
	mySQLMigrationsPath := os.Getenv("MYSQL_MIGRATIONS_PATH")
	if mySQLMigrationsPath == "" {
		mySQLMigrationsPath = config.DefaultMigrationsMysql
	}
	rawMySQLPort := os.Getenv("MYSQL_PORT")
	mySQLPort, err := strconv.Atoi(rawMySQLPort)
	if err != nil {
		return config.Config{}, fmt.Errorf("invalid MYSQL_PORT: %s", err.Error())
	}
	postgresMigrationsPath := os.Getenv("POSTGRES_MIGRATIONS_PATH")
	if postgresMigrationsPath == "" {
		postgresMigrationsPath = config.DefaultMigrationsPostgres
	}
	rawPostgresPort := os.Getenv("POSTGRES_PORT")
	postgresPort, err := strconv.Atoi(rawPostgresPort)
	if err != nil {
		return config.Config{}, fmt.Errorf("invalid POSTGRES_PORT: %s", err.Error())
	}

	// Load env variables
	cfg := config.Config{
		Server: config.Server{
			Host:           os.Getenv("HOST"),
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
		Environment: os.Getenv("ENVIRONMENT"),
		Service:     os.Getenv("SERVICE"),
	}
	return cfg, nil
}
