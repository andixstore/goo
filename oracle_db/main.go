package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

func main() {

	db, err := sql.Open(
		"godror", "plnadmin/plnadmin@apkt_dev")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select kode_asset from ss_jaringan where rownum < 10")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
		fmt.Printf("The date is: %s\n", thedate)
	}
}
