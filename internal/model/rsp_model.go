package model

type Images struct {
	Images []Image `json:"images"`
}

type Image struct {
	EndDate   string `json:"enddate"`
	Url       string `json:"url"`
	Copyright string `json:"copyright"`
}
