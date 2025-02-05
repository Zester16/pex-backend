package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"pex.oschmid.com/database"
	"pex.oschmid.com/model"
)

// type SplashInterface interface {
// 	GetAllSplash
// }

// func (database database.DBSplash) AddSplash(model.Splash) (int, error) {
// 	res, err := database.Query("INSERT into SPLASH(NAME DATE) VALUES ($1, $2)", model.Splash.Name, model.Splash.Date)
// }

func GetSplash() ([]model.SplashModel, error) {

	row, err := database.DBSplash.Query("SELECT * FROM splash")

	if err != nil {
		return nil, err

	}
	defer row.Close()
	splashes := []model.SplashModel{}
	for row.Next() {
		var splashModel model.SplashModel
		err := row.Scan(&splashModel.Id, &splashModel.Name, &splashModel.Date)

		if err != nil {
			fmt.Println("errorRepository", err)
			return []model.SplashModel{}, err
		}
		splashes = append(splashes, splashModel)
	}
	return splashes, nil
}

func GetIndividualSplash(id int) (model.SplashModel, error) {
	row, err := database.DBSplash.Query("SELECT * FROM splash WHERE id=$1", id)
	if err != nil {
		fmt.Println("repository-splash-gis", err)
		return model.SplashModel{}, errors.New("no data is present")
	}

	defer row.Close()
	var splash model.SplashModel
	var errN error
	for row.Next() {
		fmt.Println("into row")
		switch err := row.Scan(&splash.Id, &splash.Name, &splash.Date); err {
		case sql.ErrNoRows:
			log.Println("No rows were returned!")
			errN = errors.New("no data is present")
		case nil:
			log.Println("in splash", splash)
		default:
			//   panic(err)
			log.Println("Returning default!")
			errN = errors.New("no data is present")
		}
	}
	return splash, errN
}
