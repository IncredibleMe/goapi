package models

type Thema struct {
	ID      uint       `json:"id" gorm:"primary_key"`
	Onoma   string     `json:"onoma"`
	Patriko *Thema     `json:"patriko"`
	Status  StatusName `json:"status"`
}
