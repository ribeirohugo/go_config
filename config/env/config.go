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
	mongoDBPort, err := getNumber("MONGODB_PORT")
	if err != nil {
		return config.Config{}, err
	}
	mySQLMigrationsPath := os.Getenv("MYSQL_MIGRATIONS_PATH")
	if mySQLMigrationsPath == "" {
		mySQLMigrationsPath = config.DefaultMigrationsMysql
	}
	mySQLPort, err := getNumber("MYSQL_PORT")
	if err != nil {
		return config.Config{}, err
	}
	postgresMigrationsPath := os.Getenv("POSTGRES_MIGRATIONS_PATH")
	if postgresMigrationsPath == "" {
		postgresMigrationsPath = config.DefaultMigrationsPostgres
	}
	postgresPort, err := getNumber("POSTGRES_PORT")
	if err != nil {
		return config.Config{}, err
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

func getNumber(key string) (int, error) {
	rawIntValue := os.Getenv(key)
	intValue, err := strconv.Atoi(rawIntValue)
	if err != nil {
		return 0, fmt.Errorf("invalid %s: %s", key, err.Error())
	}
	return intValue, nil
}
