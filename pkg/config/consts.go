package config

const (
	DefaultMigrationsMongo    = "file://migrations/mongo"
	DefaultMigrationsMysql    = "file://migrations/mysql"
	DefaultMigrationsPostgres = "file://migrations/postgres"
	DefaultMongoPort          = 27017
	DefaultMySQLPort          = 3306
	DefaultPostgresPort       = 5432
	DefaultSessionMaxAge      = 86400 // 24 hours
	DefaultJaegerHost         = "http://localhost:14268/api/traces"
	DefaultLokiHost           = "http://localhost:3100/loki/api/v1/push"
	DefaultTempoHost          = "http://localhost:4318/v1/traces"
	DefaultRedisHost          = "localhost:6379"
)
