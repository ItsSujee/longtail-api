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

	if err := models.DB.Raw("WITH arrivals AS ( SELECT arrival.bs_id, arrival.br_id, arrival.arrival_time, arrival.mins_to_arrival,counter FROM ( SELECT bs_id, bus_runs.br_id, arrival_time :: time, floor(extract (epoch FROM arrival_time :: time) / 60 - extract(epoch FROM CURRENT_TIME(0):: time at time zone 'UTC+3') / 60) AS mins_to_arrival, row_number() OVER(partition BY br_id ORDER BY arrival_time) AS counter FROM bus_run_schedule JOIN bus_runs ON bus_run_schedule.run_id = bus_runs.run_id AND bus_runs.day_of_week = CASE WHEN extract(isodow FROM date (now() :: date)) = 7 THEN 'sunday' WHEN extract(isodow FROM date (now() :: date)) = 6 THEN 'sunday' ELSE 'weekday' END AND floor(extract (epoch FROM arrival_time :: time) / 60 - extract(epoch FROM CURRENT_TIME(0):: time at time zone 'UTC+3') / 60) > 0 AND bs_id = " + strconv.Itoa(bs_id) + ") AS arrival WHERE counter < 5 AND counter > 0) SELECT bus_routes.br_id, bus_routes.route_name, bus_routes.start_bs_id, bus_routes.end_bs_id,string_agg(arrivals.arrival_time::text, ', ') as arrival_time, string_agg(arrivals.mins_to_arrival::text, ', ') as mins_to_arrival FROM bus_routes JOIN arrivals ON bus_routes.br_id = arrivals.br_id JOIN bus_route_stops ON bus_routes.br_id = bus_route_stops.br_id WHERE bus_route_stops.bs_id = " + strconv.Itoa(bs_id) + "GROUP BY bus_routes.br_id, bus_routes.route_name, bus_routes.start_bs_id, bus_routes.end_bs_id ORDER BY br_id, mins_to_arrival ASC").Scan(&output).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Database Error:": err})
		return
	}
	c.JSON(http.StatusOK, output)

}
