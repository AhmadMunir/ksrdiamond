package models

import (
	_ "database/sql"
	"fmt"
	"kasir/db"
)

type Employee struct {
	Id     int16  `json:"id"`
	Name   string `json:"employee_name"`
	Salary string `json:"employee_salary"`
	Age    string `json:"employee_age"`
}

type Employees struct {
	Employees []Employee `json:"employee"`
}

func GetEmployee() Employees {
	con := db.CreateCon()
	//db.Createon()
	sqlStatement := "SELECT id, employee_name, employee_age, employee_salary FROM employee order by id"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated u)
	}

	defer rows.Close()
	result := Employees{}

	for rows.Next() {
		employee := Employee{}

		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
		//exit if we get an error
		if err2 != nil {
			fmt.Println(err2)
		}

		result.Employees = append(result.Employees, employee)
	}
	return result
}
