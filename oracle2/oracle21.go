package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "gopkg.in/goracle.v2"
)

func main() {
	//connect to database
	// db, err := sql.Open("goracle", "hr/hr@192.168.56.18:1521/rokoko")
	db, err := sql.Open("goracle", "reportapkt/reportapkt@apkt_dev")

	//close database connection at the end of program execution
	defer db.Close()

	//check if there were some errors during connection
	if err != nil {
		log.Panic(err)
	}

	//execute a query and get pointer to the rows
	rows, err := db.Query("select count(1) from appuser")

	//check if there were some errors during execution
	if err != nil {
		// log.Panic(err)
	}

	//define variable for the output
	var emps_cnt int

	//if there are some rows - put the output to the variable address by reference
	if rows.Next() {
		rows.Scan(&emps_cnt)
	}

	fmt.Println(emps_cnt)

}
