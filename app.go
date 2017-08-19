package main

import (
	"campoos-id/model/database/queries"
)

func main() {
	//insert new student
	//firstname,lastname,username,password,status
	queries.CreateNewStudent("Olgha", "Haq", "olghahaq", "cinta", "inactive")
	//get all students
	queries.GetAllStudents()
}
