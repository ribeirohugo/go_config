package xml

import (
	"encoding/xml"
	"io"
	"os"

	"github.com/ribeirohugo/go_config/v2/pkg/config"
)

// Load loads configurations from a given XML file path.
func Load(filePath string) (config.XML, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return config.XML{}, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return config.XML{}, err
	}
	_ = file.Close()

	return LoadContent(bytes)
}

// LoadContent loads configurations from a given xml bytes content.
func LoadContent(content []byte) (config.XML, error) {
	cfg := config.XML{
		Config: config.Config{
			MySql: config.Database{
				Port:           config.DefaultMySQLPort,
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
			Tempo: config.ExternalService{
				Host: config.DefaultTempoHost,
			},
			Token: config.Token{
				MaxAge: config.DefaultSessionMaxAge,
			},
			Jaeger: config.ExternalService{
				Host: config.DefaultJaegerHost,
			},
		},
	}

	err := xml.Unmarshal(content, &cfg)
	if err != nil {
		return config.XML{}, err
	}

	return cfg, nil
}
