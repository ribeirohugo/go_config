package config

import (
	"fmt"
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

const (
	defaultMigrationsMongo    = "file://migrations/mongo"
	defaultMigrationsMysql    = "file://migrations/mysql"
	defaultMigrationsPostgres = "file://migrations/postgres"
	defaultSessionMaxAge      = 86400 // 24 hours

	defaultJaegerHost = "http://localhost:14268/api/traces"
)

// Database holds database connection configurations.
type Database struct {
	Host           string `toml:"host"`
	Port           int    `toml:"port"`
	User           string `toml:"user"`
	Password       string `toml:"password"`
	Db             string `toml:"database"`
	MigrationsPath string `toml:"migrations_path"`
}

// Server holds server host and port configurations.
type Server struct {
	Host           string   `toml:"host"`
	Port           int      `toml:"port"`
	AllowedOrigins []string `toml:"allowed_origins"`
}

type Token struct {
	MaxAge int    `toml:"max_age"`
	Secret string `toml:"secret"`
}

// Tracer holds jaeger tracer config attributes
type Tracer struct {
	Enabled    bool   `toml:"enabled"`
	JaegerHost string `toml:"jaeger_host"`
}

// Config holds configurations data and methods.
type Config struct {
	Server Server `toml:"server"`
	Token  Token  `toml:"token"`

	MongoDb  Database `toml:"mongodb"`
	MySql    Database `toml:"mysql"`
	Postgres Database `toml:"postgres"`

	Tracer Tracer `toml:"tracer"`

	Environment string `toml:"environment"`
	Service     string `toml:"service"`
}

// Load loads configurations from a given toml file path.
func Load(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}
	_ = file.Close()

	config := Config{
		MySql: Database{
			MigrationsPath: defaultMigrationsMysql,
		},
		MongoDb: Database{
			MigrationsPath: defaultMigrationsMongo,
		},
		Postgres: Database{
			MigrationsPath: defaultMigrationsPostgres,
		},
		Token: Token{
			MaxAge: defaultSessionMaxAge,
		},
		Tracer: Tracer{
			JaegerHost: defaultJaegerHost,
		},
	}

	err = toml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

// GetAddress returns website address.
func (s Server) GetAddress() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

// MongodbAddress returns MongoDB connection address.
func (c Config) MongodbAddress() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin&ssl=false",
		c.MongoDb.User, c.MongoDb.Password, c.MongoDb.Host, c.MongoDb.Port, c.MongoDb.Db)
}

// MysqlAddress returns MySQL connection address.
func (c Config) MysqlAddress() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		c.MySql.User, c.MySql.Password, c.MySql.Host, c.MySql.Port, c.MySql.Db)
}

// PostgresAddress returns PostgreSQL connection address.
func (c Config) PostgresAddress() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.Postgres.User, c.Postgres.Password, c.Postgres.Host, c.Postgres.Port, c.Postgres.Db)
}
