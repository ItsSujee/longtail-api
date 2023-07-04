package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itssujee/longtail-api/models"
)

// GET /map_markers lat? long? deltalat? deltalong?
// Find a bus stops on the current view
func GetMapMarkers(c *gin.Context) {
	var output []models.BusStop
	lat, err := strconv.ParseFloat(c.Query("lat"), 64)
	long, err := strconv.ParseFloat(c.Query("long"), 64)
	deltalat, err := strconv.ParseFloat(c.Query("deltalat"), 64)
	deltalong, err := strconv.ParseFloat(c.Query("deltalong"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Parameter Parsing Error:": err})
		return
	}
	maxLat := lat + deltalat
	minLat := lat - deltalat
	maxLong := long + deltalong
	minLong := long - deltalong
	if err := models.DB.Where("latitude > ? AND latitude < ? AND longitude > ? AND longitude < ?", minLat, maxLat, minLong, maxLong).Find(&output).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Database Error:": "Something went wrong!"})
		return
	}
	c.JSON(http.StatusOK, output)
}
