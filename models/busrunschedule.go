package models

import (
    "time"
)

type BusRunSchedule struct {
	BrsId  uint  `json:"brs_id" gorm:"primary_key"`
	BrID  uint `json:"br_id"`
	BsId uint `json:"bs_id"`
	ArrivalTime time.Time `json:"arrival_time"`
	MinsToArrival uint `json:"mins_to_arrival"`
}