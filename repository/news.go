package repository

import (
	"pex.oschmid.com/database"
	"pex.oschmid.com/model"
)

func AddNewspaper(newspaper model.NewspaperModel) error {

	_, err := database.DBSplash.Query("INSERT INTO newspaper(id,name,created_at,image_url,epaper_url) VALUES ($1,$2,$3,$4,$5)", newspaper.Id, newspaper.Name, newspaper.Created_At, newspaper.Image_Url, newspaper.Epaper_Url)
	return err
}

func AddNewsRead(newsread model.NewspaperreadingModel) error {
	_, err := database.DBSplash.Query("INSERT INTO newsread(id,read_at,newspaper_id) VALUES ($1,$2,$3)", newsread.Id, newsread.Read_At, newsread.Newspaper_Id)
	return err
}

func GetNewspaperPaginated(id string) ([]model.NewspaperModel, error) {

	row, err := database.DBSplash.Query("SELECT id,name,image_url FROM newspaper where id > $1 ORDER BY id LIMIT 2", id)

	if err != nil {
		return nil, err
	}

	defer row.Close()
	newspapersList := []model.NewspaperModel{}
	for row.Next() {
		var newspaper model.NewspaperModel

		err := row.Scan(&newspaper.Id, &newspaper.Name, &newspaper.Image_Url)

		if err != nil {
			return []model.NewspaperModel{}, err
		}
		newspapersList = append(newspapersList, newspaper)
	}

	return newspapersList, nil
}

func GetNewspaperTotalCount() (int, error) {
	row, err := database.DBSplash.Query("SELECT COUNT(*) from newspaper")
	total := 0
	for row.Next() {

		if err != nil {
			return 0, err
		}

		row.Scan(&total)
	}
	return total, nil
}
