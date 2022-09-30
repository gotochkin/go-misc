// Copyright 2022 Gleb Otochkin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
/// CRUD operations using go-ora driver https://github.com/sijms/go-ora
package db_conn

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func getEmployeesOra(db *sql.DB) (EmpData, error) {
	employees := EmpData{}
	t1 := time.Now()
	rows, err := db.Query("SELECT employee_id,first_name,last_name,hire_date,manager_id FROM employees ORDER BY 4 DESC FETCH FIRST 10 ROWS ONLY")
	if err != nil {
		return employees, fmt.Errorf("an employees rows scan error: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			employee_id int64
			first_name  string
			last_name   string
			hire_date   time.Time
			manager_id  int64
		)
		err := rows.Scan(&employee_id, &first_name, &last_name, &hire_date, &manager_id)
		if err != nil {
			return employees, fmt.Errorf("an employees rows scan error: %v", err)
		}
		employees.Employees = append(employees.Employees, Employee{Employee_id: employee_id, First_name: first_name, Last_name: last_name, Hire_date: hire_date, Manager_id: manager_id})
	}
	//
	employees.GetResponseTime = time.Since(t1).String()
	return employees, nil
}
func PostEmployeeOra(db *sql.DB) error {
	// Filling the table with sample dataset
	var err error
	type employee struct {
		Employee_id int64
		First_name  string
		Last_name   string
		Hire_date   string
		Manager_id  int64
	}
	data := []employee{}
	data = append(data, employee{Employee_id: 100, First_name: "Steven", Last_name: "King", Hire_date: "06-17-2003", Manager_id: 100})
	data = append(data, employee{Employee_id: 101, First_name: "Neena", Last_name: "Kochhar", Hire_date: "09-21-2005", Manager_id: 100})
	data = append(data, employee{Employee_id: 102, First_name: "Lex", Last_name: "De Haan", Hire_date: "01-13-2001", Manager_id: 102})
	data = append(data, employee{Employee_id: 103, First_name: "Alexander", Last_name: "Hunold", Hire_date: "01-03-2006", Manager_id: 102})
	data = append(data, employee{Employee_id: 104, First_name: "Bruce", Last_name: "Ernst", Hire_date: "05-21-2007", Manager_id: 103})
	data = append(data, employee{Employee_id: 105, First_name: "David", Last_name: "Austin", Hire_date: "06-25-2005", Manager_id: 103})
	data = append(data, employee{Employee_id: 106, First_name: "Valli", Last_name: "Pataballa", Hire_date: "02-05-2006", Manager_id: 103})
	for i := range data {
		// Insert tyhe data
		employee_id := data[i].Employee_id
		first_name := data[i].First_name
		last_name := data[i].Last_name
		hire_date := data[i].Hire_date
		manager_id := data[i].Manager_id
		//Insert data
		insertEmp := "INSERT INTO employees(employee_id, first_name, last_name, hire_date, manager_id) VALUES( :1, :2, :3, TO_DATE(:4,'mm-dd-yyyy'), :5)"
		_, err := db.Exec(insertEmp, employee_id, first_name, last_name, hire_date, manager_id)
		if err != nil {
			log.Printf("func PostEmployee: unable to save employee: %v", err)
		}
		// fmt.Fprintf(w, "The employee %s is successfully added!", first_name)
	}
	return err

}
