package queries

import (
	"campoos-id/helper"
	"campoos-id/model/database"
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

var DB = database.ConnectDB()

func CreateNewStudent(fn string, ln string, un string, pw string, st string) {
	//defer DB.Close()
	_, err := DB.Exec(`insert into students(firstname,lastname,username,password,status)
                    VALUES(?,?,?,?,?)`, fn, ln, un, helper.HashingPassword(pw), st)
	if err != nil {
		log.Fatal("query error")
	}
	fmt.Println("new Student Created")
}

func GetStudentById(id int8) students {
	defer DB.Close()
	row := DB.QueryRow(`select * from students where id = ?`, id)
	result := students{}
	row.Scan(&result.id, &result.firstname, &result.lastname, &result.username, &result.password, &result.status)
	return result
}

func GetStudents() []students {
	defer DB.Close()
	rows, err := DB.Query("select * from students")
	if err != nil {
		return nil
	}
	defer rows.Close()
	var result = []students{}
	for rows.Next() {
		var each = students{}
		var err = rows.Scan(&each.id, &each.firstname, &each.lastname, &each.username, &each.password, &each.status)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return result
}

func DeleteStudent(id int8) {
	//defer DB.Close()
	_, err := DB.Exec(`delete from students where id = ?`, id)
	if err != nil {
		log.Fatal("query error")
	}
	fmt.Printf("Student with ID %d has been deleted\n", id)
}
