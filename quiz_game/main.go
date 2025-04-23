package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}
	var correct_answer_count int
	for _, record := range records {
		fmt.Println("What's the answer for " + record[0] + "?")
		var user_answer string
		fmt.Scanln(&user_answer)
		if record[1] == user_answer {
			correct_answer_count++
		}
	}
	fmt.Println("You aswered", correct_answer_count, "questions correctly from a set of", len(records))

}
