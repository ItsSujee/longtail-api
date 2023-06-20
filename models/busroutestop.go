package models

type BusRouteStop struct {
	BrsId  uint  `json:"brs_id" gorm:"primary_key"`
	BrID  uint `json:"br_id"`
	BsId uint `json:"bs_id"`
	SeqNum uint `json:"seq_num"`
}