package models

import "time"

type OrderH struct {
	OrderID      int       `json:"order_id" gorm:"primary_key;auto_increment:true"`
	OrderNo      string    `json:"order_no" gorm:"type:varchar(20)"`
	BarberID     int       `json:"barber_id" gorm:"type:integer"`
	CapsterID    int       `json:"capster_id" gorm:"type:integer"`
	OrderDate    time.Time `json:"order_date" gorm:"type:timestamp(0) without time zone"`
	UserID       int       `json:"user_id" gorm:"type:integer"`
	CustomerName string    `json:"customer_name" gorm:"type:varchar(60);not null"`
	Email        string    `json:"email" gorm:"type:varchar(60)"`
	Telp         string    `json:"telp" gorm:"type:varchar(20)"`
	Status       string    `json:"status" gorm:"type:varchar(1)"`
	FromApps     bool      `json:"from_apps" gorm:"type:boolean"`
	UserInput    string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit     string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput    time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit     time.Time `json:"time_Edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

type OrderPost struct {
	OrderID    int       `json:"order_id" gorm:"type:integer"`
	BarberID   int       `json:"barber_id" gorm:"type:integer"`
	BarberName string    `json:"barber_name" gorm:"type:varchar(60)"`
	Latitude   float64   `json:"latitude" gorm:"type:float8"`
	Longitude  float64   `json:"longitude" gorm:"type:float8"`
	UserID     int       `json:"user_id,omitempty" gorm:"type:integer"`
	Name       string    `json:"name,omitempty" gorm:"type:varchar(100)"`
	Email      string    `json:"email,omitempty" gorm:"type:varchar(60)"`
	OrderDate  time.Time `json:"order_date" gorm:"type:timestamp(0) without time zone;default:now()"`
}

type OrderList struct {
	OwnerID      int       `json:"owner_id"`
	BarberID     int       `json:"barber_id" valid:"Required"`
	BarberName   string    `json:"barber_name"`
	OrderID      int       `json:"order_id"`
	OrderNo      string    `json:"order_no"`
	Status       string    `json:"status" `
	FromApps     bool      `json:"from_apps"`
	CapsterID    int       `json:"capster_id,omitempty"`
	CapsterName  string    `json:"capster_name"`
	CustomerName string    `json:"customer_name"`
	OrderDate    time.Time `json:"order_date" valid:"Required"`
	FileID       int       `json:"file_id" `
	FileName     string    `json:"file_name"`
	FilePath     string    `json:"file_path"`
	Price        float32   `json:"price" `
	// Weeks       int       `json:"weeks"`
	// Months      int       `json:"months"`
	// Years       int       `json:"years"`
}
type OrderStatus struct {
	Status string `json:"status" `
}

type OrderNotif struct {
	TimeArrive   int64     `json:"time_arrive" gorm:"time_arrive;type:integer"`
	OrderID      int       `json:"order_id" gorm:"primary_key;auto_increment:true"`
	OrderNo      string    `json:"order_no" gorm:"type:varchar(20)"`
	CapsterID    int       `json:"capster_id" gorm:"type:integer"`
	OrderDate    time.Time `json:"order_date" gorm:"type:timestamp(0) without time zone"`
	UserID       int       `json:"user_id" gorm:"type:integer"`
	CustomerName string    `json:"customer_name" gorm:"type:varchar(60);not null"`
}

func (OrderNotif) TableName() string {
	return "order_h"
}
