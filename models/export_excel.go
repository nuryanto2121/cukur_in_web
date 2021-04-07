package models

import "time"

type TrxApp struct {
	OrderNo      string    `json:"order_no" gorm:"type:varchar(20)"`
	OrderDate    time.Time `json:"order_date" gorm:"type:timestamp(0) without time zone"`
	BarberCd     string    `json:"barber_cd" gorm:"type:varchar(20)"`
	BarberName   string    `json:"barber_name" gorm:"type:varchar(20)"`
	CapsterName  string    `json:"capster_name"`
	CustomerName string    `json:"customer_name"`
	FromApps     bool      `json:"from_apps"`
	PaketName    string    `json:"paket_name"`
	Price        float32   `json:"price"`
}
