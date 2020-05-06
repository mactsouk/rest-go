package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mactsouk/handlers"
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
		fmt.Println(nil)
		return
	}

	temp := handlers.User{}
	fmt.Println(temp)

}
