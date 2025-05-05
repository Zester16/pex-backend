package repository

import (
	"pex.oschmid.com/database"
	"pex.oschmid.com/model"
)

func AddNewspaper(newspaper model.NewspaperModel) error {

	_, err := database.DBSplash.Query("INSERT INTO newspaper(id,name,created_at) VALUES ($1,$2,$3)", newspaper.Id, newspaper.Name, newspaper.Created_At)
	return err
}

func AddNewsRead(newsread model.NewspaperreadingModel) error {
	_, err := database.DBSplash.Query("INSERT INTO newsread(id,read_at,newspaper_id) VALUES ($1,$2,$3)", newsread.Id, newsread.Read_At, newsread.Newspaper_Id)
	return err
}
