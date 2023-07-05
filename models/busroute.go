package models

type BusRoute struct {
	BrId          uint   `json:"br_id" gorm:"primary_key"`
	RouteName     string `json:"route_name"`
	StartBsId     uint   `json:"start_bs_id"`
	EndBsId       uint   `json:"end_bs_id"`
	ArrivalTime   string `json:"arrival_time"`
	MinsToArrival string `json:"mins_to_arrival"`
}
