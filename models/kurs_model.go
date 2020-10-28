package models

import (
	_ "database/sql"
	"fmt"
	"kasir/db"
	kurs "kasir/domain"
)

// var con *sql.DB

func InsertKurs(rp, coin int) error {
	db := db.CreateCon()
	_, err := db.Exec("INSERT INTO smilecoin values (?,?,?)", "", rp, coin)
	if err != nil {
		fmt.Println(err.Error())
	}
	db.Close()

	fmt.Println("Insert Kurs Success!")

	return err
}

func GetKurs() (kurs.Kurs, error) {
	con := db.CreateCon()

	sqlSt := "SELECT * FROM kurs WHERE id = 1"

	row, err := con.Query(sqlSt)
	fmt.Println(row)
	if err != nil {
		fmt.Println(err)
	}

	result := kurs.Kurs{}

	for row.Next() {

		err2 := row.Scan(&result.Id, &result.Rp, &result.Coin)
		if err2 != nil {
			fmt.Print(err2)
		}

	}
	con.Close()

	return result, err
}
