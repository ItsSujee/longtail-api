package controllers

import (
"net/http"
"github.com/gin-gonic/gin"
"github.com/itssujee/longtail-api/models"
)

// GET /bus_route
// Get all bus_routes
func FindBusRoutes(c *gin.Context) {
	var GetRoutes []models.BusRoute
	models.DB.Find(&GetRoutes)
	c.JSON(http.StatusOK, GetRoutes)
}

// GET /bus_route/:br_id
// Find a bus_route
func FindBusRoute(c *gin.Context) {
	var GetRoute models.BusRoute
	if err := models.DB.Where("br_id = ?", c.Param("br_id")).First(&GetRoute).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
	c.JSON(http.StatusOK, GetRoute)
}