package model

type NewspaperModel struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Created_At *int64 `json:"created_at"`
	Image_Url string `json:"image_url"`
	Epaper_Url string `json:"epaper_url"`
}

type NewspaperreadingModel struct {
	Id           string
	Read_At      *int64
	Newspaper_Id string
}
