// Package system provides OS-level utilities: environment/config loading,
// shell command execution, file checksums, and debugger detection.
package system

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

// LoadEnvs loads environment variables from .env and .env.default files
// in the given directory.
func LoadEnvs(path string) error {
	path = path + "/"
	return godotenv.Load(path+".env", path+".env.default")
}

// ReadConfig reads a YAML configuration file, selects the given environment
// section, and decodes it into the provided pointer using mapstructure.
func ReadConfig(path string, env string, pointer interface{}) error {
	readFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var yml map[string]map[string]interface{}
	if err := yaml.Unmarshal(readFile, &yml); err != nil {
		return err
	}

	if yml[env] == nil {
		return errors.New("no environment: " + env)
	}

	return mapstructure.WeakDecode(yml[env], &pointer)
}
