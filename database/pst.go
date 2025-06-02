package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DBSplash *sql.DB

func ConnectDB() error {
	var err error
	pstUrl := os.Getenv("PST_URL")
	DBSplash, err = sql.Open("postgres", pstUrl)

	if err != nil {
		fmt.Println("database Error", err)
		return err
	}
	fmt.Println("Database Connected")

	createNewsPaperTable()
	return nil
}

// stubbing this for future use
func createNewsPaperTable() {

	_, err := DBSplash.Query(`CREATE TABLE IF NOT EXISTS newspaper (
    id text PRIMARY KEY,
    name text UNIQUE NOT NULL,
    created_at int8 UNIQUE NOT NULL
)
`)
	if err != nil {
		fmt.Println("Database creation issue", err.Error())
	}
	_, errTwo := DBSplash.Query(`CREATE TABLE IF NOT EXISTS newsread (
    id text PRIMARY KEY,
    read_at int8 UNIQUE NOT NULL,
	newspaper_id text NOT NULL,
    CONSTRAINT newspaper_id FOREIGN KEY (newspaper_id)
	REFERENCES newspaper(id)

)
`)
	if errTwo != nil {
		fmt.Println("Database creation issue", err.Error())
	}
}
