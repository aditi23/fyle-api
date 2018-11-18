package db

import (
	"database/sql"
	"fmt"
)

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

var DB *sql.DB

func InitDB() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

// dbConfig heroku config
func dbConfig() map[string]string {
	conf := make(map[string]string)
	conf[dbhost] = "localhost"
	conf[dbport] = "5432"
	conf[dbuser] = ""
	conf[dbpass] = ""
	conf[dbname] = "fyle"
	return conf
}
