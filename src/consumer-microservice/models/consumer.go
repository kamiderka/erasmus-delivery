package models

import (
	"github.com/google/uuid"
)

// User information
type Consumer struct {
	UUID     		uuid.UUID 	`json:"id" `
	FirstName   	string    	`json:"firstName"`
	LastName  		string		`json:"lastName"`
    Email			string		`json:"email"`
	Address			string		`json:"address"`
	City			string		`json:"city"`
	Password 		string		`json:"password"`
}