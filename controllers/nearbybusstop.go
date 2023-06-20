package controllers

import (
"net/http"
"github.com/gin-gonic/gin"
"github.com/itssujee/longtail-api/models"
"strconv"
)

// GET /nearby_bus_stops lat? long? n?
// Get n closest busstops
func FindNearbyBusStops(c *gin.Context) {
	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Fatal Error": "Converting Variable n to int!"})
		return
	}
	lat := c.Query("lat")
	long := c.Query("long")
	var GetNearbyBusStops []models.BusStop
	models.DB.Select("bs_id,name,cross_street1,cross_street2,latitude,longitude,SQRT( POWER(latitude - " + lat + ", 2) + POWER(longitude - " + long + ", 2) ) as distance").Order("distance").Limit(n).Find(&GetNearbyBusStops)
	c.JSON(http.StatusOK, GetNearbyBusStops)
}
