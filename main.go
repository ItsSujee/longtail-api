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
			"message": "Routes: /map_markers /nearby_bus_stops /bus_routes_from_bus_id/:bs_id /bus_run_schedule",
		})
	})

	// /bus_stops
	// r.GET("/bus_stops", controllers.FindBusStops)
	// r.GET("/bus_stops/:bs_id", controllers.FindBusStop)

	// /bus_routes
	// r.GET("/bus_routes", controllers.FindBusRoutes)
	// r.GET("/bus_routes/:br_id", controllers.FindBusRoute)

	// /bus_route_stops
	// r.GET("/bus_route_stops", controllers.FindBusRouteStops)
	// r.GET("/bus_route_stops/:brs_id", controllers.FindBusRouteStop)

	// /map_markers
	r.GET("/map_markers", controllers.GetMapMarkers)

	// /nearby_bus_stops
	r.GET("/nearby_bus_stops", controllers.FindNearbyBusStops)

	// /bus_routes_from_bus_id
	r.GET("/bus_routes_from_bus_id/:bs_id", controllers.FindBusRoutesFromStop)

	// /bus_run_schedule
	r.GET("/bus_run_schedule", controllers.FindBusSchedule)

	r.Run()
}
