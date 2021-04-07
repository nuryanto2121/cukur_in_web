package models

import "time"

type RedemTeguk struct {
	RedemID     int       `json:"redem_id" gorm:"primary_key;auto_increment:true"`
	RedemCd     string    `json:"redem_cd" gorm:"type:varchar(20)"`
	OrderID     int       `json:"order_id" gorm:"type:integer"`
	BarberID    int       `json:"barber_id" gorm:"type:integer"`
	IsUsed      bool      `json:"is_used" gorm:"type:boolean"`
	ExpiredDate time.Time `json:"expired_date" gorm:"type:timestamp(0) without time zone;default:now()"`
	Model
}
