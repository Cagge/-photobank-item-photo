package models

type PHotodate struct {
	User           string
	Pass           string
	BaseUrlListing string
	Urldownload    string
}
type ConfigDataPhoto struct {
	Photodate PHotodate
}

type Data struct {
	Count  int `json:"count"`
	Photos []Phot
}

type Phot struct {
	Number int    `json:"number"`
	Sid    int    `json:"sid"`
	Type   string `json:"type"`
}
