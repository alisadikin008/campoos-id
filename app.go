package main

import (
	"campoos-id/model/database/queries"
	"fmt"
)

func main() {
	//insert new student
	//firstname,lastname,username,password,status
	//queries.CreateNewStudent("Dyah", "Sari", "dyahsari", "april1986", "inactive")
	//get all students
	//get student by ID
	fmt.Println(queries.GetStudentById(1))
	// queries.DeleteStudent(14)
	// students := queries.GetStudents()
	// for _, student := range students {
	// 	fmt.Println(student)
	// 	//fmt.Println(student["id"], student.firstname, student.lastname, student.username, student.password, student.status)
	// }
}
