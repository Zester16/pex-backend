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

	return nil
}

//stubbing this for future use
// func CreateProductTable() {
//     DB.Query(`CREATE TABLE IF NOT EXISTS products (
//     id SERIAL PRIMARY KEY,
//     amount integer,
//     name text UNIQUE,
//     description text,
//     category text NOT NULL
// )
// `)
// }
