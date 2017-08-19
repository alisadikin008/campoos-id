package queries

import (
	"campoos-id/model/database"
	"crypto/md5"
	"fmt"
	"log"
)

type students struct {
	id        int8
	firstname string
	lastname  string
	username  string
	password  string
	status    string
}

type input_student struct {
	firstname string
	lastname  string
	username  string
	password  string
	status    string
}

var DB = database.ConnectDB()

func hashingPassword(password string) {
	hasher := md5.New()
	hasher.Write([]byte(password))
	//hex.EncodeToString(hasher.Sum(nil))
}

func CreateNewStudent(fn string, ln string, un string, pw string, st string) {
	//defer DB.Close()
	_, err := DB.Exec("insert into students VALUES(?,?,?,?,?)", fn, ln, un, hashingPassword(pw), st)
	if err != nil {
		log.Fatal("query error")
	}
	fmt.Println("new Student Created")
}

func GetAllStudents() {
	defer DB.Close()
	rows, err := DB.Query("select id,firstname,lastname from students")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var result []students
	for rows.Next() {
		var each = students{}
		var err = rows.Scan(&each.id, &each.firstname, &each.lastname)
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
	fmt.Println("ID", "Firstname", "Lastname")
	for _, each := range result {
		fmt.Println(each.id, each.firstname, each.lastname)
	}
}
