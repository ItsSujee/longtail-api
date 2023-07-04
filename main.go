package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itssujee/longtail-api/controllers"
	"github.com/itssujee/longtail-api/models"
)

var Router *gin.Engine

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// root /
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "/map_markers?lat,long,deltalat,deltalong /bus_stops?lat,long,n /bus_routes?bs_id",
		})
	})

	// /map_markers lat? long? deltalat? deltalong?
	r.GET("/map_markers", controllers.GetMapMarkers)

	// /bus_stops lat? long? n?
	r.GET("/bus_stops", controllers.GetBusStops)

	// /bus_routes bs_id?
	r.GET("/bus_routes", controllers.GetBusRoutes)

	r.Run()
}
