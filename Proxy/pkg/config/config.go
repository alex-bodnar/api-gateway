package config

import (
	"errors"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

// LoadFromFile function which load config from file via filepath.
func LoadFromFile(filepath string, dest interface{}) error {
	if t := reflect.TypeOf(dest); t.Kind() != reflect.Ptr {
		return errors.New("can`t load config on non-pointer value")
	}

	fileData, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(fileData, dest); err != nil {
		return err
	}

	return nil
}
