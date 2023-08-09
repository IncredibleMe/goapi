package models

type Thema struct {
	ID      uint       `json:"id" gorm:"primary_key"`
	Onoma   string     `json:"onoma"`
	Patriko string     `json:"patriko"`
	Status  StatusName `json:"status"`
}
