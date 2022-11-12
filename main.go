package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var students = []Student{

	{Id: 1, Name: "Burak", Class: "1-B", Teacher: "Burak Tabakoglu"},
	{Id: 2, Name: "AHMET", Class: "2-B", Teacher: "Burak Tabakoglu"},
}

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Class   string `json:"class"`
	Teacher string `json:"teacher"`
}

func listStudents(Context *gin.Context) {
	Context.IndentedJSON(http.StatusOK, students)

}

func createStudent(Context *gin.Context) {
	var studentByUser Student
	err := Context.BindJSON(&studentByUser)

	if err == nil && studentByUser.Id != 1 && studentByUser.Class != "" && studentByUser.Name != "" && studentByUser.Teacher != "" {

		students = append(students, studentByUser)
		Context.IndentedJSON(http.StatusCreated, gin.H{"mesage": "Student has been created.", "Student_id": studentByUser.Id})
		return
	} else {
		Context.IndentedJSON(http.StatusBadRequest, gin.H{"mesage": "Student cannot created.!"})
		return

	}

}

func getStudentByID(int_id int) (*Student, error) {

	for i, s := range students {
		if s.Id == int_id {
			return &students[i], nil
		}

	}

	return nil, errors.New("Student cannot be found")
}

func getStudent(Context *gin.Context) {

	str_id := Context.Param("id")
	int_id, err := strconv.Atoi(str_id)

	if err != nil {
		panic(err)
	}

	student, err := getStudentByID(int_id)
	if err == nil {
		Context.IndentedJSON(http.StatusOK, student)
	} else {

		Context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Student cannot be found!!"})
	}

}

func main() {

	router := gin.Default()
	router.GET("/students", listStudents)
	router.POST("/students", createStudent)
	router.GET("/students/:id", getStudent)
	router.Run("localhost:9090")

}
