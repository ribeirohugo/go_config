package dotenv

import (
	"github.com/joho/godotenv"

	"github.com/ribeirohugo/go_config/v2/pkg/config"
	"github.com/ribeirohugo/go_config/v2/pkg/config/env"
)

// Load loads configurations from a given toml file path.
func Load(filePath string) (config.Config, error) {
	err := godotenv.Load(filePath)
	if err != nil {
		return config.Config{}, err
	}

	return env.Load()
}
