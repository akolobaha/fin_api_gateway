package entities

type Securities []Security

type Security struct {
	Ticker    string `gorm:"ticker"`
	Shortname string `gorm:"shortname"`
	Secname   string `gorm:"secname"`
}
