package config

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

// Tracer holds jaeger tracer config attributes
type Tracer struct {
	Enabled    bool   `toml:"enabled" yaml:"enabled"`
	JaegerHost string `toml:"jaeger_host" yaml:"jaeger_host"`
}
