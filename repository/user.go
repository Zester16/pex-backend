package repository

import (
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
		fmt.Println("Session-Check", err)
		return model.SessionModel{}, errors.New("some error in dB")
	}
	defer row.Close()
	if row.Next() {
		//for row.Next() {
		if err := row.Scan(&session.Id, &session.Created_At, &session.Device, &session.Username, &session.Last_Login, &session.Expiry); err != nil {
			fmt.Println("Error when scanning get session", err.Error())
			return model.SessionModel{}, err
		}
		// case sql.ErrNoRows:
		// 	err = errors.New("no rows were returned")

		// case nil:
		// 	fmt.Println("successful transaction", session.Expiry)
		// default:
		// 	err = errors.New("no data is present")
		// }

		//}
	} else {
		err = errors.New("no data is present")
	}

	return session, err
}

// get all sessions paginated, FE should pass the last id
// func GetAllSessionsWithPagination(prevSid string) ([]model.SessionModel, error) {

// }

// update last login time
func UpdateLoginTimeForSession(id string, lastLogin int64) {

	_, err := database.DBSplash.Query(`UPDATE session SET last_login=$1 WHERE id=$2`, lastLogin, id)
	if err != nil {
		fmt.Println(err)
	}

}

func DeleteSession(id string) error {
	succ, err := database.DBSplash.Query("DELETE from session WHERE id=$1", id)

	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("success", succ)

	return err
}

func GetAllSessions() ([]model.SessionModel, error) {

	allSession := []model.SessionModel{}

	row, err := database.DBSplash.Query("SELECT ALL * FROM session")
	if err != nil {
		fmt.Println("Session-Check", err)
		return []model.SessionModel{}, errors.New("some error in dB")
	}
	defer row.Close()

	// row.Next() {
	var session model.SessionModel
	//session.Last_Login = int64(0)
	for row.Next() {
		if err := row.Scan(&session.Id, &session.Created_At, &session.Device, &session.Username, &session.Last_Login, &session.Expiry); err != nil {
			return allSession, err
		}
		// case sql.ErrNoRows:
		// 	err = errors.New("no rows were returned")

		// case nil:
		// 	fmt.Println("successful transaction", session.Expiry)
		// default:
		// 	err = errors.New("no data is present")
		// }

		allSession = append(allSession, session)

	}
	if err = row.Err(); err != nil {
		return allSession, err
	}
	// } else {
	// 	err = errors.New("no data is present")
	//}

	return allSession, err

}
