package controllers

import (
"net/http"
"github.com/gin-gonic/gin"
"github.com/itssujee/longtail-api/models"
)

// GET /bus_route_stop
// Get all bus_route_stops
func FindBusRouteStops(c *gin.Context) {
	var GetRoutes []models.BusRouteStop
	models.DB.Find(&GetRoutes)
	c.JSON(http.StatusOK, GetRoutes)
}

// GET /bus_route_stops/:brs_id
// Find a bus_route_stop by id
func FindBusRouteStop(c *gin.Context) {
	var GetRoute models.BusRouteStop
	if err := models.DB.Where("brs_id = ?", c.Param("brs_id")).First(&GetRoute).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
	c.JSON(http.StatusOK, GetRoute)
  }
