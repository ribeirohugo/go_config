package yaml

import (
	"io"
	"os"

	"github.com/ribeirohugo/go_config/config"

	"gopkg.in/yaml.v3"
)

const (
	defaultMigrationsMongo    = "file://migrations/mongo"
	defaultMigrationsMysql    = "file://migrations/mysql"
	defaultMigrationsPostgres = "file://migrations/postgres"
	defaultSessionMaxAge      = 86400 // 24 hours

	defaultJaegerHost = "http://localhost:14268/api/traces"
)

// Load loads configurations from a given yaml file path.
func Load(filePath string) (config.Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return config.Config{}, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return config.Config{}, err
	}
	_ = file.Close()

	cfg := config.Config{
		MySql: config.Database{
			MigrationsPath: defaultMigrationsMysql,
		},
		MongoDb: config.Database{
			MigrationsPath: defaultMigrationsMongo,
		},
		Postgres: config.Database{
			MigrationsPath: defaultMigrationsPostgres,
		},
		Token: config.Token{
			MaxAge: defaultSessionMaxAge,
		},
		Tracer: config.Tracer{
			JaegerHost: defaultJaegerHost,
		},
	}

	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}
