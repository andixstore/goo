package main

import (
	"database/sql"
	"log"

	_ "github.com/godror/godror"
	// _ "gopkg.in/goracle.v2"
)

var (
	dbusername = "reportapkt"
	dbpassword = "reportapkt"
	dbsid      = "apkt_dev"
)

/*
ConnectDB connects to the database and returns a sql.DB pointer
*/
func ConnectDB() *sql.DB {

	// connString := fmt.Sprintf("%s/%s/@//****:****/%s",
	// dbusername,
	// dbpassword,
	// dbsid)

	// db, err := sql.Open("goracle", connString)
	db, err := sql.Open(
		"godror", "plnadmin/plnadmin@apkt_dev")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
