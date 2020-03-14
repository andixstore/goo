package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-oci8"
)

func main() {
	db, err := sql.Open("oci8", "user/password@XE")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	loginStmt, err := db.Prepare("begin user_mgt.login(:1,:2,:3);end;")
	defer loginStmt.Close()
	mail := "ofoe"
	pwd := "testing"
	var retVal string
	var retStatus string
	er := loginStmt.QueryRow(mail, pwd, retVal).Scan(&retStatus)
	if er != nil {
		panic(er)
	}
	fmt.Println("retval")
	fmt.Println(retStatus)
}
