package config

const (
	DefaultMigrationsMongo    = "file://migrations/mongo"
	DefaultMigrationsMysql    = "file://migrations/mysql"
	DefaultMigrationsPostgres = "file://migrations/postgres"
	DefaultSessionMaxAge      = 86400 // 24 hours

	DefaultJaegerHost = "http://localhost:14268/api/traces"
)
