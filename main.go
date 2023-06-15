package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itssujee/longtail-api/models"
	"github.com/itssujee/longtail-api/controllers"
)

var Router * gin.Engine

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// root /
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	// /bus_stops
	r.GET("/bus_stops", controllers.FindBusStops)
	r.GET("/bus_stops/:bs_id", controllers.FindBusStop)

	r.Run()
}
