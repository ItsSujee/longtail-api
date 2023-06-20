package controllers

import (
"net/http"
"github.com/gin-gonic/gin"
"github.com/itssujee/longtail-api/models"
"strconv"
)

// GET /bus_stop
// Get all bus_stops
func FindBusStops(c *gin.Context) {
	var GetStops []models.BusStop
	models.DB.Find(&GetStops)
	c.JSON(http.StatusOK, GetStops)
}

// GET /bus_stop/:bs_id
// Find a bus_stop
func FindBusStop(c *gin.Context) {
	var GetStop models.BusStop
	if err := models.DB.Where("bs_id = ?", c.Param("bs_id")).First(&GetStop).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
	c.JSON(http.StatusOK, GetStop)
  }


// GET /map_markers/:lat long deltalat deltalong
// Find a bus_stop based on lat long deltalat deltalong 
func GetMapMarkers(c *gin.Context) {
	var GetStops []models.BusStop
	lat,err := strconv.ParseFloat(c.Query("lat"), 64)
	long,err := strconv.ParseFloat(c.Query("long"), 64)
	deltalat,err := strconv.ParseFloat(c.Query("deltalat"), 64)
	deltalong, err := strconv.ParseFloat(c.Query("deltalong"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"API error:": "Float Parsing Error"})	
		return
	}
	maxLat := lat + deltalat
	minLat := lat - deltalat
	maxLong := long + deltalong
	minLong := long - deltalong
	if err := models.DB.Where("latitude > ? AND latitude < ? AND longitude > ? AND longitude < ?", minLat, maxLat, minLong, maxLong ).Find(&GetStops).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"API error:": "Something went wrong!"})
	  return
	}
	c.JSON(http.StatusOK, GetStops)
  }