package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
)

var questions = []string{}

func takeQfile() error {
	fd, err := os.Open("problems.csv")
	if err != nil {
		return err
	}

	file_reader := csv.NewReader(fd)
	records, err := file_reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file content", err)
	}
	fmt.Println("Questions read in")

	fmt.Println(reflect.TypeOf(records))

	// var questions = []string{}

	for _, i := range records {
		// qst := strings.Fields(i[0])
		// questions = append(questions, qst)
		fmt.Println(reflect.TypeOf(i))
	}
	defer fd.Close()

	return nil
}

func QuizMaster() string {
	return ""
}

func CheckAnswer() string {
	return ""
}

func main() {
	takeQfile()
	fmt.Println("hey gugugaga")

}
