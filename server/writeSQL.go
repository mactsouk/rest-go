package main

import (
	"database/sql"
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need SQLite3 Database File")
		return
	}
	database := arguments[1]

	fmt.Println("Writing to SQLite3:", database)
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		fmt.Println(nil)
		return
	}

	fmt.Println("Emptying database table.")
	_, err = db.Exec("DELETE FROM data")
	if err != nil {
		fmt.Println(nil)
		return
	}

}
