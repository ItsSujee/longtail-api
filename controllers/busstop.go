package controllers

import (
"net/http"
"github.com/gin-gonic/gin"
"github.com/itssujee/longtail-api/models"
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
