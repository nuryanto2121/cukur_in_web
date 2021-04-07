package models

import "time"

// import uuid "github.com/satori/go.uuid"

type SsUser struct {
	UserID      int       `json:"user_id" gorm:"PRIMARY_KEY"`
	Name        string    `json:"name" gorm:"type:varchar(60);not null"`
	Telp        string    `json:"telp" gorm:"type:varchar(20)"`
	Email       string    `json:"email" gorm:"type:varchar(60)"`
	IsActive    bool      `json:"is_active" gorm:"type:boolean"`
	JoinDate    time.Time `json:"join_date" gorm:"type:timestamp(0);default:now()"`
	BirthOfDate time.Time `json:"birth_of_date" gorm:"type:timestamp(0)"`
	Password    string    `json:"password" gorm:"type:varchar(150)"`
	FileID      int       `json:"file_id" gorm:"type:integer"`
	UserType    string    `json:"user_type" gorm:"type:varchar(10)"`
	UserInput   string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit    string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput   time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit    time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

type AddUser struct {
	Email    string `json:"email" valid:"Required"`
	Telp     string `json:"telp"`
	Password string `json:"password"`
	Name     string `json:"name" valid:"Required"`
	IsAdmin  bool   `json:"is_admin"`
}

type UpdateUser struct {
	Email string `json:"email" valid:"Required"`
	Telp  string `json:"telp"`
	Name  string `json:"name" valid:"Required"`
}

type LoginCapster struct {
	CapsterID      int    `json:"capster_id"`
	CapsterName    string `json:"capster_name"`
	IsActive       bool   `json:"is_active"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Telp           string `json:"telp"`
	FileID         int    `json:"file_id"`
	FileName       string `json:"file_name"`
	FilePath       string `json:"file_path"`
	BarberID       int    `json:"barber_id"`
	BarberName     string `json:"barber_name"`
	BarberIsActive bool   `json:"barber_is_active"`
	OwnerID        int    `json:"owner_id"`
	OwnerName      string `json:"owner_name"`
}
