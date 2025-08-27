package model

// ***DB MODELS**/
// newspaper model for DB
type NewspaperModel struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Created_At *int64 `json:"created_at"`
	Total_Read int    `json:"total_read"`
	Last_Read  *int64 `json:"last_read"`
	Image_Url  string `json:"image_url"`
	Epaper_Url string `json:"epaper_url"`
}

type NewspaperreadingModel struct {
	Id           string
	Read_At      *int64 `json:"read_at"`
	Newspaper_Id string `json:"newspaper_id"`
}

// FRONTENT MODELS
// To respond back to front end
type NewsFEResponseModel struct {
	Id           string
	Read_At      *int64 `json:"read_at"`
	Newspaper_Id string `json:"newspaper_id"`
	Image_Url    string `json:"image_url"`
	Name         string `json:"name"`
}
