package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init() {
	const dbString string = "brothers.db"
	var err error
	db, err = sql.Open("sqlite3", dbString)

	if err != nil {
		log.Fatal(err)
	}

	// defer db.Close()

	// var version string
	// err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(version)
}

func GetDB() *sql.DB {
	return db
}
