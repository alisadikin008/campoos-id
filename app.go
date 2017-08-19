package main

import (
	"campoos-id/model/database/queries"
)

func main() {
	//insert new student
	//firstname,lastname,username,password,status
	queries.CreateNewStudent("Dyah", "Sari", "dyahsari", "april1986", "inactive")
	//get all students
	queries.GetAllStudents()
}
