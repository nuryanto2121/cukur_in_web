package models

type PatnerMaster struct {
	CompanionId int `json:"companion_id" gorm:"primary_key;auto_increment:true"`
	AddPatnerMaster
	Model
}

type AddPatnerMaster struct {
	Name      string  `json:"name" gorm:"type:varchar(60)"`
	Address   string  `json:"address" gorm:"type:varchar(150)"`
	GroupName string  `json:"group_name" gorm:"type:varchar(60)"`
	Latitude  float64 `json:"latitude" gorm:"type:float8"`
	Longitude float64 `json:"longitude" gorm:"type:float8"`
}
