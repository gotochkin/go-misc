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
	"log"
)

func checkDBObjectOra(db *sql.DB, objname string) (int, error) {
	//
	//defer elapsedTime(time.Now(), "chekObject")
	var cnt int
	err := db.QueryRow("select count(*) from user_objects where object_name = :1", objname).Scan(&cnt)
	if err != nil {
		//return 0, fmt.Errorf("DB.QueryRow: %v", err)
		return -1, err
	}
	if cnt > 0 {
		return cnt, nil
	}
	return cnt, nil
}
func execStmt(tdll string) error {
	//
	_, err := db.Exec(tdll)
	return err
}

func runDDL(db *sql.DB) error {
	//Create table if it doesn't exist
	var errddl error
	createEmp := `CREATE TABLE employees (
		employee_id NUMBER(6) NOT NULL,
		first_name VARCHAR(20) NOT NULL,
		last_name VARCHAR(20) NOT NULL,
		hire_date DATE NOT NULL,
		manager_id NUMBER(6),
		PRIMARY KEY (employee_id)
	)`

	//Create the employee table
	errddl = execStmt(createEmp)
	if errddl != nil {
		log.Fatalf("Unable to create object: %s", errddl)
	}
	return errddl
}
