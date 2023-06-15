package models

type BusStop struct {
	BsId  uint  `json:"bs_id" gorm:"primary_key"`
	Name  string `json:"name"`
	CrossStreet1 string `json:"cross_street1"`
	CrossStreet2 string `json:"cross_street2"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}