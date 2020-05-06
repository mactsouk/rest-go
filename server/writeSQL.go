package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
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
	_, _ = db.Exec("DROP TABLE data")

	fmt.Println("Creating table from scratch.")
	_, err = db.Exec("CREATE TABLE data (Username STRING, Password STRING, Admin Bool);")
	if err != nil {
		fmt.Println(nil)
		return
	}

	fmt.Println("Populating", database)
	stmt, _ := db.Prepare("INSERT INTO data(Username, Password, Admin) values(?,?,?)")
	for i := 20; i < 50; i++ {
		if i%2 == 0 {
			_, _ = stmt.Exec("name"+strconv.Itoa(i), "pass"+strconv.Itoa(2*i), true)
		} else {
			_, _ = stmt.Exec("name"+strconv.Itoa(i), "pass"+strconv.Itoa(2*i), false)
		}
	}
}
