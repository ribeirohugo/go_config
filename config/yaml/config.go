package yaml

import (
	"io"
	"os"

	"github.com/ribeirohugo/go_config/config"

	"gopkg.in/yaml.v3"
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

	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}
