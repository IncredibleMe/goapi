package models

type Movie struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	Titlos         string `json:"titlos"`
	Periexomeno    string `json:"periexomeno"`
	Thema          Thema  `json:"thema"`
	Hm_dimiourgias string `json:"im_dimiourgias"`
}
