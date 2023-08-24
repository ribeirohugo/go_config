package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// LoadYaml loads configurations from a given yaml file path.
func LoadYaml(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}
	_ = file.Close()

	config := Config{
		MySql: Database{
			MigrationsPath: defaultMigrationsMysql,
		},
		MongoDb: Database{
			MigrationsPath: defaultMigrationsMongo,
		},
		Postgres: Database{
			MigrationsPath: defaultMigrationsPostgres,
		},
		Token: Token{
			MaxAge: defaultSessionMaxAge,
		},
		Tracer: Tracer{
			JaegerHost: defaultJaegerHost,
		},
	}

	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
