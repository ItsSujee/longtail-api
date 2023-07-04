package models

import (
    "time"
)

type BusRunSchedule struct {
	BrsId  uint  `json:"brs_id" gorm:"primary_key"`
	ArrivalTime time.Time `json:"arrival_time"`
	MinsToArrival uint `json:"mins_to_arrival"`
}