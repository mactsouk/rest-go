package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mactsouk/handlers"
	_ "github.com/mattn/go-sqlite3"
)

const (
	empty = ""
	tab   = "\t"
)

func PrettyJson(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}

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

	var c1, c2 string
	var c3 bool

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3)
		temp := handlers.Input{c1, c2, c3}
		t, _ := PrettyJson(temp)
		fmt.Println(t)
	}
}
