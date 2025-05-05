package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"strings"
)

var db *sql.DB

func main() {

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASSWD")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "tasks"
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	perr := db.Ping()
	if perr != nil {
		fmt.Print(perr.Error())
	}

	if len(os.Args) == 1 {
		fmt.Println("Usage:")
		fmt.Println("\ttask [command]")
		fmt.Println("Available commands")
		fmt.Println("\tadd Add a new task to your TODO list")
		fmt.Println("\tdo Mark a task on your TODO list as complete")
		fmt.Println("\tlist List all of your incomplete tasks")
		os.Exit(1)
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) == 2 {
			panic("Command add should have a parameter")
		}
		query := fmt.Sprintf("INSERT INTO task (description) VALUES ('%s')", strings.Join(os.Args[2:], " "))
		_, err := db.Query(query)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Error inserting task")
		}
	case "do":
		if len(os.Args) != 3 {
			panic("Command do should have a parameter")
		}
		task_id := os.Args[2]
		query := "SELECT * FROM task WHERE id = " + task_id
		row := db.QueryRow(query)
		var id int
		var description string
		if err != nil {
			panic("Error deleting task")
		}
		err := row.Scan(&id, &description)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("No rows with this id in the database")
			os.Exit(1)
		case nil:

		default:
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if _, err := db.Exec("DELETE FROM task WHERE id = " + task_id); err != nil {
			fmt.Println("Error deleting task")
		}
		fmt.Println("You have completed the \"" + description + "\" task")

	case "list":
		if len(os.Args) > 2 {
			panic("Task list doesn't take any parameters")
		}
		fmt.Println("You have the following tasks:")
		query := "SELECT * FROM task"
		rows, err := db.Query(query)
		defer rows.Close()
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Error listing tasks")
		}
		for rows.Next() {
			var id int
			var description string
			err := rows.Scan(&id, &description)
			if err != nil {
				fmt.Println("Error scaning row in list command")
			}
			fmt.Println(strconv.Itoa(id) + ". " + description)
		}
	default:
		panic("Command " + command + " does not exist")
	}

}
