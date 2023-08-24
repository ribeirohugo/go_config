package config

import "fmt"

// Config holds configurations data and methods.
type Config struct {
	Server Server `toml:"server" yaml:"server" json:"server,omitempty"`
	Token  Token  `toml:"token" yaml:"token" json:"token,omitempty"`

	MongoDb  Database `toml:"mongodb" yaml:"mongodb" json:"mongodb,omitempty"`
	MySql    Database `toml:"mysql" yaml:"mysql" json:"mysql,omitempty"`
	Postgres Database `toml:"postgres" yaml:"postgres" json:"postgres,omitempty"`

	Tracer Tracer `toml:"tracer" yaml:"tracer" json:"tracer,omitempty"`

	Environment string `toml:"environment" yaml:"environment" json:"environment,omitempty"`
	Service     string `toml:"service" yaml:"service" json:"service,omitempty"`
}

// Database holds database connection configurations.
type Database struct {
	Host           string `toml:"host" yaml:"host" json:"host,omitempty"`
	Port           int    `toml:"port" yaml:"port" json:"port,omitempty"`
	User           string `toml:"user" yaml:"user" json:"user,omitempty"`
	Password       string `toml:"password" yaml:"password" json:"password,omitempty"`
	Db             string `toml:"database" yaml:"database" json:"database,omitempty"`
	MigrationsPath string `toml:"migrations_path" yaml:"migrations_path" json:"migrations_path,omitempty"`
}

// Server holds server host and port configurations.
type Server struct {
	Host           string   `toml:"host" yaml:"host" json:"host,omitempty"`
	Port           int      `toml:"port" yaml:"port" json:"port,omitempty"`
	AllowedOrigins []string `toml:"allowed_origins" yaml:"allowed_origins" json:"allowed_origins,omitempty"`
}

// Token holds application token secret and expire time in seconds.
type Token struct {
	MaxAge int    `toml:"max_age" yaml:"max_age" json:"max_age,omitempty"`
	Secret string `toml:"secret" yaml:"secret" json:"secret,omitempty"`
}

// Tracer holds jaeger tracer toml attributes
type Tracer struct {
	Enabled    bool   `toml:"enabled" yaml:"enabled" json:"enabled,omitempty"`
	JaegerHost string `toml:"jaeger_host" yaml:"jaeger_host" json:"jaeger_host,omitempty"`
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
