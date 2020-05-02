package main

import (
	"database/sql"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Reading from SQLite3")

	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need SQLite3 Database File")
		return
	}
	database := arguments[1]

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		fmt.Println(nil)
		return
	}
}
