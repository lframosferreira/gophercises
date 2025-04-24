package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var correct_answer_count int

func ask_questions(records *[][]string) {
	for _, record := range *records {
		question := record[0]
		answer := record[1]
		fmt.Println("What's the answer for " + question + "?")
		var user_answer string
		fmt.Scanln(&user_answer)
		if answer == user_answer {
			correct_answer_count++
		}
	}
}

func main() {

	problems_file_name := flag.String("file", "problems.csv", "The name of the problems csv file")
	time_limit := flag.Int("time_limit", 30, "Time limit for quiz")
	shuffle_problems := flag.Bool("shuffle_problems", true, "Should the input list be shuffled?")
	flag.Parse()
	file, err := os.Open(*problems_file_name)
	if err != nil {
		panic(err)
	}
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}
	if *shuffle_problems {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })
	}
	go func() {
		time.Sleep(time.Duration(*time_limit) * time.Second)
		fmt.Println("You aswered", correct_answer_count, "questions correctly from a set of", len(records))
		os.Exit(0)
	}()
	ask_questions(&records)

	fmt.Println("You aswered", correct_answer_count, "questions correctly from a set of", len(records))

}
