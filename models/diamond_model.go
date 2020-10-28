package models

import (
	"database/sql"
	_ "database/sql"
	"fmt"

	// "github.com/AhmadMunir/kasir/db/db_diamond.go"
	"kasir/db"

	dm "kasir/domain"
)

var con *sql.DB

func PanicError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func InsertDiamond(name string, dm string, price int) error {
	db := db.CreateCon()
	defer db.Close()

	_, err := db.Exec("insert into tbl_diamond values (?, ?, ?, ?)", "", name, dm, price)
	if err != nil {
		fmt.Println(err.Error())
		// return
		PanicError(err)
	}

	fmt.Println("Insert Success!")

	return err
}

func GetAllDiamonds() (dm.Diamonds, error) {

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM tbl_diamond order by id_diamond"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	result := dm.Diamonds{}

	for rows.Next() {
		diamond := dm.Diamond{}

		err2 := rows.Scan(&diamond.Id, &diamond.Name, &diamond.Dm, &diamond.Price)
		if err2 != nil {
			fmt.Println(err2)
		}

		result.Diamonds = append(result.Diamonds, diamond)
	}

	return result, err
}

func DeleteDiamondById(id int) error {
	con := db.CreateCon()

	_, err := con.Exec("DELETE from tbl_diamond WHERE id_diamond =?", id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	con.Close()
	fmt.Println("Delete Success!")
	return err
}

func UpdateDiamondById(id, price int, name, dm string) error {
	con := db.CreateCon()

	_, err := con.Exec("UPDATE tbl_diamond SET name = ?, diamond = ?, price = ? WHERE id_diamond = ?", name, dm, price, id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Data Success!")

	return err
}
