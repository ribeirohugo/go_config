package xml

import (
	"encoding/xml"
	"io"
	"os"

	"github.com/ribeirohugo/go_config/pkg/config"
)

// Load loads configurations from a given XML file path.
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

// LoadContent loads configurations from a given xml bytes content.
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
		Loki: config.ExternalService{
			Host: config.DefaultLokiHost,
		},
		Token: config.Token{
			MaxAge: config.DefaultSessionMaxAge,
		},
		Jaeger: config.ExternalService{
			Host: config.DefaultJaegerHost,
		},
	}

	err := xml.Unmarshal(content, &cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}
