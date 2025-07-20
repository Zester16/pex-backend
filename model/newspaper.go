package model

type NewspaperModel struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Created_At *int64 `json:"created_at"`
	Total_Read int `json:"total_read"`
	Last_Read *int64 `json:"last_read"`
	Image_Url string `json:"image_url"`
	Epaper_Url string `json:"epaper_url"`
}

type NewspaperreadingModel struct {
	Id           string
	Read_At      *int64
	Newspaper_Id string
}
