package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itssujee/longtail-api/models"
)

// GET /bus_stops lat? long? n?
// Get n closest bus stops
func GetBusStops(c *gin.Context) {
	var output []models.BusStop
	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Parameter Parsing Error": err})
		return
	}
	lat := c.Query("lat")
	long := c.Query("long")
	if err := models.DB.Select("bs_id,name,cross_street1,cross_street2,latitude,longitude,SQRT( POWER(latitude - " + lat + ", 2) + POWER(longitude - " + long + ", 2) ) as distance").Order("distance").Limit(n).Find(&output).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Database Error:": err})
		return
	}
	c.JSON(http.StatusOK, output)
}
