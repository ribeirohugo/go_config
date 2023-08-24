package config

import "fmt"

// Config holds configurations data and methods.
type Config struct {
	Server Server `toml:"server" yaml:"server"`
	Token  Token  `toml:"token" yaml:"token"`

	MongoDb  Database `toml:"mongodb" yaml:"mongodb"`
	MySql    Database `toml:"mysql" yaml:"mysql"`
	Postgres Database `toml:"postgres" yaml:"postgres"`

	Tracer Tracer `toml:"tracer" yaml:"tracer"`

	Environment string `toml:"environment" yaml:"environment"`
	Service     string `toml:"service" yaml:"service"`
}

// Database holds database connection configurations.
type Database struct {
	Host           string `toml:"host" yaml:"host"`
	Port           int    `toml:"port" yaml:"port"`
	User           string `toml:"user" yaml:"user"`
	Password       string `toml:"password" yaml:"password"`
	Db             string `toml:"database" yaml:"database"`
	MigrationsPath string `toml:"migrations_path" yaml:"migrations_path"`
}

// Server holds server host and port configurations.
type Server struct {
	Host           string   `toml:"host" yaml:"host"`
	Port           int      `toml:"port" yaml:"port"`
	AllowedOrigins []string `toml:"allowed_origins" yaml:"allowed_origins"`
}

// Token holds application token secret and expire time in seconds.
type Token struct {
	MaxAge int    `toml:"max_age" yaml:"max_age"`
	Secret string `toml:"secret" yaml:"secret"`
}

// Tracer holds jaeger tracer toml attributes
type Tracer struct {
	Enabled    bool   `toml:"enabled" yaml:"enabled"`
	JaegerHost string `toml:"jaeger_host" yaml:"jaeger_host"`
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
