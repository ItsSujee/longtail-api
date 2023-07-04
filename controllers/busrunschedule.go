package controllers

import (
"net/http"
"github.com/gin-gonic/gin"
"github.com/itssujee/longtail-api/models"
"strconv"
)

// GET /bus_run_schedule bs_id? br_id?
// Get bus schedule from bus stop id and bus route id
func FindBusSchedule(c *gin.Context) {
	bs_id,err := strconv.Atoi(c.Query("bs_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Fatal Error": "bs_id is not an int"})
		return
	}
	br_id,err := strconv.Atoi(c.Query("br_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Fatal Error": "br_id is not an int"})
		return
	}
	var schedule []models.BusRunSchedule
	models.DB.Raw("SELECT brs_id,bs_id, bus_runs.br_id, arrival_time::time, FLOOR(EXTRACT (epoch from arrival_time::time) / 60 - EXTRACT(epoch from CURRENT_TIME(0)::time at time zone 'UTC+3') / 60) as mins_to_arrival FROM bus_run_schedule JOIN bus_runs on bus_run_schedule.run_id = bus_runs.run_id WHERE bs_id = " + strconv.Itoa(bs_id) + " AND br_id = " + strconv.Itoa(br_id) + " AND bus_runs.day_of_week = CASE WHEN EXTRACT(isodow from date (NOW()::date)) = 7 THEN 'sunday' WHEN EXTRACT(isodow from date (NOW()::date)) = 6 THEN 'sunday' ELSE 'weekday' END AND FLOOR(EXTRACT (epoch from arrival_time::time) / 60 - EXTRACT(epoch from CURRENT_TIME(0)::time at time zone 'UTC+3') / 60) > 0 ORDER BY mins_to_arrival asc LIMIT 5;").Scan(&schedule)
	c.JSON(http.StatusOK, schedule)
}
