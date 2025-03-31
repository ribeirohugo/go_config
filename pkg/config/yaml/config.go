package yaml

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/ribeirohugo/go_config/v2/pkg/config"
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

	return LoadContent(bytes)
}

// LoadContent loads configurations from a given yaml bytes content.
func LoadContent(content []byte) (config.Config, error) {
	cfg := config.Config{
		MySql: config.Database{
			Port:           config.DefaultMySqlPort,
			MigrationsPath: config.DefaultMigrationsMysql,
		},
		MongoDb: config.Database{
			Port:           config.DefaultMongoPort,
			MigrationsPath: config.DefaultMigrationsMongo,
		},
		Postgres: config.Database{
			Port:           config.DefaultPostgresPort,
			MigrationsPath: config.DefaultMigrationsPostgres,
		},
		Token: config.Token{
			MaxAge: config.DefaultSessionMaxAge,
		},
		Loki: config.ExternalService{
			Host: config.DefaultLokiHost,
		},
		Tempo: config.ExternalService{
			Host: config.DefaultTempoHost,
		},
		Jaeger: config.ExternalService{
			Host: config.DefaultJaegerHost,
		},
	}

	err := yaml.Unmarshal(content, &cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}
