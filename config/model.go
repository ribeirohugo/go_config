package config

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

// Token holds application token secret and expire time in seconds.
type Token struct {
	MaxAge int    `toml:"max_age"`
	Secret string `toml:"secret"`
}

// Tracer holds jaeger tracer config attributes
type Tracer struct {
	Enabled    bool   `toml:"enabled"`
	JaegerHost string `toml:"jaeger_host"`
}
