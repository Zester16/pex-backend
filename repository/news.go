package repository

import (
	"fmt"

	"pex.oschmid.com/database"
	"pex.oschmid.com/model"
)

// add newspaper
func AddNewspaper(newspaper model.NewspaperModel) error {

	_, err := database.DBSplash.Query("INSERT INTO newspaper(id,name,created_at,image_url,epaper_url) VALUES ($1,$2,$3,$4,$5)", newspaper.Id, newspaper.Name, newspaper.Created_At, newspaper.Image_Url, newspaper.Epaper_Url)
	return err
}

// add newsread
func AddNewsRead(newsread model.NewspaperreadingModel) error {
	_, err := database.DBSplash.Query("INSERT INTO newsread(id,read_at,newspaper_id) VALUES ($1,$2,$3)", newsread.Id, newsread.Read_At, newsread.Newspaper_Id)
	return err
}

// add new last read date and total count
func UpdateNewspaperLastReadAndTodaysDate(id string, readDate *int64) error {
	newspaper, err := GetNewspaperById(id)

	if err != nil {

		return err
	}
	readCount := newspaper.Total_Read + 1
	lastRead := newspaper.Last_Read
	if lastRead == nil {
		lastRead = readDate
	} else if *readDate > *newspaper.Last_Read {
		lastRead = readDate
	}

	fmt.Println("UpdateNewspaperLastReadAndTodaysDate", lastRead, readCount)

	_, err = database.DBSplash.Query("UPDATE newspaper SET total_read=$1, last_read=$2 WHERE id=$3", readCount, &lastRead, id)

	return err
}

// get count of news read by pass ing two parameters and see if newspaper has been read(without using read_status parameter)
func GetNewsreadwithDateAndName(newspaperId string, read_at *int64) ([]model.NewspaperreadingModel, error) {
	row, err := database.DBSplash.Query("SELECT id, read_at, newspaper_id FROM newsread WHERE newspaper_id=$1 AND read_at=$2", newspaperId, read_at)

	var newsreadArray = []model.NewspaperreadingModel{}
	if err != nil {
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		var newsRead model.NewspaperreadingModel

		err := row.Scan(&newsRead.Id, &newsRead.Read_At, &newsRead.Newspaper_Id)

		if err != nil {
			fmt.Println(".repository.GetNewsreadwithDateAndTime.error", err.Error())
			return nil, err
		}

		newsreadArray = append(newsreadArray, newsRead)
	}

	return newsreadArray, err
}

// *************Get Newspaper by id*********************************************
func GetNewspaperById(id string) (model.NewspaperModel, error) {

	row, err := database.DBSplash.Query("SELECT id,name,total_read,last_read FROM newspaper WHERE id=$1", id)

	FUNCTION_NAME := "GetNewspaperById"
	if err != nil {
		fmt.Println(FUNCTION_NAME, " error:", err.Error())
		return model.NewspaperModel{}, err
	}

	var newspaper model.NewspaperModel

	for row.Next() {
		err := row.Scan(&newspaper.Id, &newspaper.Name, &newspaper.Total_Read, &newspaper.Last_Read)

		if err != nil {
			return newspaper, err
		}
	}
	return newspaper, nil

}

// ********** To Get Paginated newspaper list*********************************
// first iteration does not need to pass id, but further iterations need to pass id
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

// ***********To get all newspapers list***************
func GetNewspapersAll() ([]model.NewspaperModel, error) {

	row, err := database.DBSplash.Query("SELECT id,name,image_url FROM newspaper")

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

// ********** To get total newspaper count***************
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

func GetNewsReadAll() ([]model.NewsFEResponseModel, error) {

	row, err := database.DBSplash.Query("SELECT newsread.id, newsread.read_at, newsread.newspaper_id, newspaper.image_url, newspaper.name FROM newsread INNER JOIN newspaper ON newsread.newspaper_id = newspaper.id ORDER BY newsread.read_at DESC")

	if err != nil {
		return nil, err
	}
	defer row.Close()

	newsReadList := []model.NewsFEResponseModel{}
	for row.Next() {
		var newsread model.NewsFEResponseModel
		err := row.Scan(&newsread.Id, &newsread.Read_At, &newsread.Newspaper_Id, &newsread.Image_Url, &newsread.Name)

		if err != nil {
			fmt.Printf("repository.GetNewsRead error: ", err.Error())
			return nil, err
		}
		newsReadList = append(newsReadList, newsread)
	}
	return newsReadList, nil
}
