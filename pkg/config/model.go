package config

import (
	"encoding/xml"
	"fmt"
)

// Config holds configurations data and methods.
type Config struct {
	Server Server `toml:"server" yaml:"server" json:"server,omitempty" xml:"server"`
	Token  Token  `toml:"token" yaml:"token" json:"token,omitempty" xml:"token"`

	MongoDb  Database `toml:"mongodb" yaml:"mongodb" json:"mongodb,omitempty" xml:"mongodb"`
	MySql    Database `toml:"mysql" yaml:"mysql" json:"mysql,omitempty" xml:"mysql"` //nolint:revive
	Postgres Database `toml:"postgres" yaml:"postgres" json:"postgres,omitempty" xml:"postgres"`

	Audit      ExternalService `toml:"audit" yaml:"audit" json:"audit,omitempty" xml:"audit"`
	Jaeger     ExternalService `toml:"jaeger" yaml:"jaeger" json:"jaeger,omitempty" xml:"jaeger"`
	Loki       ExternalService `toml:"loki" yaml:"loki" json:"loki,omitempty" xml:"loki"`
	Tempo      ExternalService `toml:"tempo" yaml:"tempo" json:"tempo,omitempty" xml:"tempo"`
	Prometheus ExternalService `toml:"prometheus" yaml:"prometheus" json:"prometheus,omitempty" xml:"prometheus"`
	Redis      ExternalService `toml:"redis" yaml:"redis" json:"redis,omitempty" xml:"redis"`

	Environment string `toml:"environment" yaml:"environment" json:"environment,omitempty" xml:"environment"`
	Service     string `toml:"service" yaml:"service" json:"service,omitempty" xml:"service"`
}

// XML holds configurations data and methods, with XML support.
type XML struct {
	Config

	XMLName xml.Name `xml:"config"`
}

// Database holds database connection configurations.
type Database struct {
	Host           string `toml:"host" yaml:"host" json:"host,omitempty" xml:"host"`
	Port           int    `toml:"port" yaml:"port" json:"port,omitempty" xml:"port"`
	User           string `toml:"user" yaml:"user" json:"user,omitempty" xml:"user"`
	Password       string `toml:"password" yaml:"password" json:"password,omitempty" xml:"password"`
	Db             string `toml:"database" yaml:"database" json:"database,omitempty" xml:"database"`
	MigrationsPath string `toml:"migrations_path" yaml:"migrations_path" json:"migrations_path,omitempty" xml:"migrations_path"`
}

// Server holds server host and port configurations.
type Server struct {
	Host           string   `toml:"host" yaml:"host" json:"host,omitempty" xml:"host"`
	Port           int      `toml:"port" yaml:"port" json:"port,omitempty" xml:"port"`
	AllowedOrigins []string `toml:"allowed_origins" yaml:"allowed_origins" json:"allowed_origins,omitempty" xml:"allowed_origins"` //nolint:lll
}

// Token holds application token secret and expire time in seconds.
type Token struct {
	MaxAge int    `toml:"max_age" yaml:"max_age" json:"max_age,omitempty" xml:"max_age"`
	Secret string `toml:"secret" yaml:"secret" json:"secret,omitempty" xml:"secret"`
}

// ExternalService holds essential external service configuration data.
type ExternalService struct {
	Enabled bool   `toml:"enabled" yaml:"enabled" json:"enabled,omitempty" xml:"enabled"`
	Host    string `toml:"host" yaml:"host" json:"host,omitempty" xml:"host"`
	Token   string `toml:"token" yaml:"token" json:"token,omitempty" xml:"token"`
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
