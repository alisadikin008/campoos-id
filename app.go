package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_HOST = "tcp(127.0.0.1:3306)"
	DB_NAME = "db_belajar_golang"
	DB_USER = "root"
	DB_PASS = "vagrant"
	DB_TYPE = "mysql"
)

type student struct {
	id    int
	name  string
	age   int
	grade int
}

func connect() *sql.DB {
	db, err := sql.Open(DB_TYPE, DB_USER+":"+DB_PASS+"@"+DB_HOST+"/"+DB_NAME)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	db := connect()
	defer db.Close()
	rows, err := db.Query("select * from tb_student")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var result []student
	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.age, &each.grade)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ID", "Name", "Age", "Grade")
	for _, each := range result {
		fmt.Println(each.id, each.name, each.age, each.grade)
	}
}
