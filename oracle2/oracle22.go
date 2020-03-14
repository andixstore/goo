package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "gopkg.in/goracle.v2"
)

func main() {
	//connect to database
	db, err := sql.Open("goracle", "hr/hr@192.168.56.18:1521/rokoko")

	//close database connection at the end of program execution
	defer db.Close()

	//check if there were some errors during connection
	if err != nil {
		log.Panic(err)
	}

	//execute a query and get pointer to the rows
	deptId := 20
	rows, err := db.Query("select first_name, last_name, salary from employees where department_id=:1", deptId)

	//check if there were some errors during execution
	if err != nil {
		log.Panic(err)
	}

	//define variable for the output
	var salary int
	var fname, lname string

	//if there are some rows - put the output to the variable address by reference
	for rows.Next() {
		rows.Scan(&fname, &lname, &salary)
		fmt.Println(fname, lname, salary)
	}

}
