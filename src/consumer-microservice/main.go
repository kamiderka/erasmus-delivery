package main

import (
	//"consumer-microservice/utils"
	"consumer-microservice/common"
	"consumer-microservice/dtos"
    
	"github.com/gin-gonic/gin"
	"net/http"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error

	// Load config file
	err = common.LoadConfig()
	if err != nil {
		return err
	}

	// // Initialize User database
	// err = databases.Database.Init()
	// if err != nil {
	// 	return err
	// }

	m.router = gin.Default()

	return nil
}



func main() {
	m := Main{}

    if m.initServer() != nil {
		return
	}

	m.router.GET("/", func(c *gin.Context) {

		dto := dtos.RegisterConsumerDTO{
			FirstName: "Kamil",
			LastName:  "Kordecki",
			Email:     "test@o2.pl",
			Password:  "Password123",
		}
	
		
		
		errs := dtos.ValidateDTO(dto)

		if len(errs) > 0 {
			errors := []string{}
			for _, e := range errs {
				errors = append(errors, e.Error())
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
		
		
		c.JSON(http.StatusOK, gin.H{"message": "Validation successful!"})
		})

	m.router.Run(common.Config.Port)
}