package json

import (
	"encoding/json"
	"io"
	"os"

	"github.com/ribeirohugo/go_config/config"
)

// Load loads configurations from a given json file path.
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
			MigrationsPath: config.DefaultMigrationsMysql,
		},
		MongoDb: config.Database{
			MigrationsPath: config.DefaultMigrationsMongo,
		},
		Postgres: config.Database{
			MigrationsPath: config.DefaultMigrationsPostgres,
		},
		Token: config.Token{
			MaxAge: config.DefaultSessionMaxAge,
		},
		Tracer: config.Tracer{
			JaegerHost: config.DefaultJaegerHost,
		},
	}

	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

// LoadContent loads configurations from a given json bytes content.
func LoadContent(content []byte) (config.Config, error) {
	cfg := config.Config{
		MySql: config.Database{
			MigrationsPath: config.DefaultMigrationsMysql,
		},
		MongoDb: config.Database{
			MigrationsPath: config.DefaultMigrationsMongo,
		},
		Postgres: config.Database{
			MigrationsPath: config.DefaultMigrationsPostgres,
		},
		Token: config.Token{
			MaxAge: config.DefaultSessionMaxAge,
		},
		Tracer: config.Tracer{
			JaegerHost: config.DefaultJaegerHost,
		},
	}

	err := json.Unmarshal(content, &cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}
