package controllers

import (
"net/http"
"github.com/gin-gonic/gin"
"github.com/itssujee/longtail-api/models"
)

// GET /bus_routes_from_bus_id/:bs_id
// Get bus routes from bus stop id
func FindBusRoutesFromStop(c *gin.Context) {
	var n []models.BusRoute
	models.DB.Where( "br_id IN (?)", models.DB.Table("bus_route_stops").Select("br_id").Where("bs_id = ?", c.Param("bs_id"))).Find(&n)
	c.JSON(http.StatusOK, n)
}
