package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itssujee/longtail-api/models"
)

// GET /bus_routes bs_id?
// Get bus routes from bus stop id
func GetBusRoutes(c *gin.Context) {
	var output []models.BusRoute
	bs_id, err := strconv.Atoi(c.Query("bs_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Parameter Parsing Error": err})
		return
	}
	if err := models.DB.Raw("WITH arrivals as ( SELECT arrival.bs_id, arrival.br_id, arrival.arrival_time, arrival.mins_to_arrival FROM ( SELECT bs_id, bus_runs.br_id, arrival_time::time,  FLOOR(EXTRACT (epoch from arrival_time::time) / 60 - EXTRACT(epoch from CURRENT_TIME(0)::time at time zone 'UTC+3') / 60) as mins_to_arrival, ROW_NUMBER() OVER(PARTITION BY br_id ORDER BY arrival_time) AS counter FROM bus_run_schedule  JOIN bus_runs on bus_run_schedule.run_id = bus_runs.run_id AND bus_runs.day_of_week = CASE WHEN EXTRACT(isodow from date (NOW()::date)) = 7 THEN 'sunday' WHEN EXTRACT(isodow from date (NOW()::date)) = 6 THEN 'sunday' ELSE 'weekday' END AND FLOOR(EXTRACT (epoch from arrival_time::time) / 60 - EXTRACT(epoch from CURRENT_TIME(0)::time at time zone 'UTC+3') / 60) > 0 AND bs_id = " + strconv.Itoa(bs_id) + " ORDER BY br_id, mins_to_arrival ) as arrival WHERE counter < 5) SELECT  bus_routes.br_id, bus_routes.route_name, bus_routes.start_bs_id,  bus_routes.end_bs_id, arrivals.arrival_time, arrivals.mins_to_arrival FROM bus_routes JOIN arrivals ON bus_routes.br_id = arrivals.br_id WHERE arrivals.bs_id = " + strconv.Itoa(bs_id) + " ORDER BY br_id, mins_to_arrival asc ").Scan(&output).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Database Error:": err})
		return
	}
	c.JSON(http.StatusOK, output)
}
