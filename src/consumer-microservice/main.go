package main

import (
	"consumer-microservice/common"
    
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
		c.String(http.StatusOK, common.Config.Port)
	})

	m.router.Run(common.Config.Port)
}