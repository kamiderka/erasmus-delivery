package common

import (
	"encoding/json"
	"os"

)

// Configuration stores setting values
type Configuration struct {
	Port         string `json:"port"`

	PgAddrs      string `json:"PgAddrs"`
	PgDbName     string `json:"PgDbName"`
	PgDbUsername string `json:"PgDbUsername"`
	PgDbPassword string `json:"PgDbPassword"`
}

// Config shares the global configuration
var (
	Config *Configuration
)

func LoadConfig() error {
	// Filename is the path to the json config file
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}

	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	return nil
}
