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
/// DDL and DDL operations are prformed for a samle table with employees data

package db_conn

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
)

func Usage() {
	fmt.Println()
	fmt.Println("Database CRUD operations")
	fmt.Println("  sample code with DDL/DML operations on Oracle database")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println(`  -h hostname -p port -s service -u username -pw password -w wallet_path`)
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println(`  crud_ops -h adb.us-ashburn-1.oraclecloud.com -p 1522 -s m7c5hdup4eqqydh_glebatp02_tp.adb.oraclecloud.com -u myuser -pw MyStrongPassword -w /wallet`)
	fmt.Println()
}

var (
	db   *sql.DB
	data EmpData
	err  error
)

type Employee struct {
	Employee_id int64
	First_name  string
	Last_name   string
	Hire_date   time.Time
	Manager_id  int64
}
type EmpData struct {
	Employees       []Employee
	GetResponseTime string
}

// Database Pool
func configurePool(db *sql.DB) {

	// Maximum number of connections in idle connection pool.
	db.SetMaxIdleConns(3)

	// Maximum number of open connections to the database.
	db.SetMaxOpenConns(10)

	// Maximum time (in seconds) that a connection can remain open.
	db.SetConnMaxLifetime(1800 * time.Second)

}
func RunApp(dbUser string, dbPwd string, dbHost string, dbPort int, dbName string, dbWallet string) {
	//Connect to the database
	db, err = connectOracle(dbUser, dbPwd, dbHost, dbPort, dbName, dbWallet)
	if err != nil {
		log.Fatalf("connectOracle: failed to create connection: %v", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Can't ping the database ", err)
	}
	fmt.Println("INFO: Connected!")

	// Check if the table exists in the database and create if it not there

	chk, chkerr := checkDBObjectOra(db, "EMPLOYEES")
	if chkerr != nil {
		log.Fatalf("checkDBObjectOra: Unable to verify the object - %s", chkerr)
	}
	if chk == 0 {
		//
		errddl := runDDL(db)
		if errddl != nil {
			log.Fatalf("runDDL: Unable to create object: %s", err)
		}
	} else {
		fmt.Println("INFO: The object with the same name is found in the database.")
	}
	data, err = getEmployeesOra(db)
	if len(data.Employees) == 0 {
		fmt.Printf("INFO: getEmployeesOra: The table employees is empty and will be filled by data")
		err = PostEmployeeOra(db)
		if err != nil {
			fmt.Printf("ERR: PostEmployeeOra: cannot insert data %s", err)
		}
	} else {
		fmt.Printf("INFO: getEmployeesOra: The table has been filled already and contains %s rows\n", strconv.Itoa(len(data.Employees)))
	}
	//

}
