package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"pex.oschmid.com/database"
	"pex.oschmid.com/model"
)

// create a new session
func AddSession(userSession model.SessionModel) error {

	_, err := database.DBSplash.Query("INSERT into session(id,username,device,created_at,expiry) VALUES ($1,$2,$3,$4,$5) ", userSession.Id, userSession.Username, userSession.Device, userSession.Created_At, userSession.Expiry)

	return err
}

// get a particular session
func GetSession(sId string) (model.SessionModel, error) {
	var session model.SessionModel
	var err error
	row, err := database.DBSplash.Query("SELECT * FROM session WHERE id=$1", sId)

	if err != nil {
		return model.SessionModel{}, errors.New("Some error in DB")
	}
	defer row.Close()

	for row.Next() {
		switch err := row.Scan(&session.Id, &session.Created_At, &session.Expiry, &session.Device, &session.Last_Login, &session.Username); err {

		case sql.ErrNoRows:
			err = errors.New("no rows were returned")

		case nil:
			fmt.Println("successful transaction", session.Expiry)
		default:
			err = errors.New("no data is present")
		}

	}
	return session, err
}

// update last login time
func UpdateLoginTimeForSession(id string, lastLogin int64) {

	_, err := database.DBSplash.Query(`UPDATE session SET last_login=$1 WHERE id=$2`, lastLogin, id)
	if err != nil {
		fmt.Println(err)
	}

}

// func DeleteSession(id string) error {
// 	succ, err := database.DBSplash.Query("DELETE ")

// 	return
// }
