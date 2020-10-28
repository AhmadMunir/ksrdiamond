package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//create mysql connection
func CreateCon() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/diamond")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is Connected")
	}

	//defer db.Close()
	//make sure connection is avaible
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("ERROR! DB is not Connected")
		fmt.Println(err.Error())
	}

	return db
}

//end mysql connection
