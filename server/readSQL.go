package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mactsouk/handlers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need SQLite3 Database File")
		return
	}
	database := arguments[1]

	fmt.Println("Reading from SQLite3:", database)
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		fmt.Println(nil)
		return
	}

	var c1 string
	var c2 string
	var c3 bool

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3)
		temp := handlers.Input{c1, c2, c3}
		fmt.Printf("%v\n", temp)
	}

}
