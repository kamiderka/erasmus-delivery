package common

import (
	"encoding/json"
	"os"
)

// Configuration stores setting values
type Configuration struct {
	Port         		string `json:"port"`

	PgAddrs      		string `json:"PgAddrs"`
	PgDbName     		string `json:"PgDbName"`
	PgDbUsername 		string `json:"PgDbUsername"`
	PgDbPassword 		string `json:"PgDbPassword"`

	JwtSecretPassword 	string `json:"jwtSecretPassword"`
	Issuer            	string `json:"issuer"`

}

// Config shares the global configuration
var (
	Config *Configuration
)

// Status Text
const (
	ErrNameEmpty      		= "Name fields is empty"
	ErrPasswordEmpty  		= "Password fields is empty"
	ErrPasswordDoNotMatch 	= "Password fields do not match"
	ErrEmailEmpty 			= "Email field is empty"
	ErrInvalidEmail 		= "Email is invalid"
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
