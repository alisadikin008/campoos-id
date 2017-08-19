package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const DB_TYPE = "mysql"
const DB_USER = "root"
const DB_PASSWORD = "vagrant"
const DB_HOST = "tcp(127.0.0.1:3306)"
const DB_NAME = "campoos_id"

func ConnectDB() *sql.DB {
	//https://stackoverflow.com/questions/15698479/how-to-connect-to-mysql-with-go
	db, err := sql.Open(DB_TYPE, DB_USER+":"+DB_PASSWORD+"@"+DB_HOST+"/"+DB_NAME)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
