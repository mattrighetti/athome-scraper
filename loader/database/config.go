package database

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Migrations  string `yaml:"migrations"`
	UpdateQuery string `yaml:"updateQuery"`
	InsertQuery string `yaml:"insertQuery"`
}

func ParseConfig(filepath string) Configuration {
	yamlBytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	config := Configuration{}
	err = yaml.Unmarshal(yamlBytes, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
