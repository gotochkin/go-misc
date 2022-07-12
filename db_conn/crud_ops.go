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

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	go_ora "github.com/sijms/go-ora/v2"
)

func usage() {
	fmt.Println()
	fmt.Println("CRUD operations")
	fmt.Println("  sample code with DDL/DML operations with Oracle database")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println(`  crud_ops -cs server_url`)
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println(`  crud_ops -cs "oracle://user:pass@server/service_name"`)
	fmt.Println()
}

var (
	conn_str = flag.String("cs", "", "Connection URL, oracle://user:pass@server/service_name")
)

func main() {
	flag.Parse()
	fmt.Println(*conn_str)
	urlOptions := map[string]string{
		"trace file": "trace.log",
	}
	// databaseURL := go_ora.BuildUrl(server, port, service, user, password, urlOptions)
	// conn, err := sql.Open("oracle", databaseURL)
	// // check error
	databaseURL := go_ora.BuildUrl("host", "1522", "m5c5hpup7eqqydh_glebatp02_tp.adb.oraclecloud.com", "admin", "ShittyPassword01#", urlOptions)
	fmt.Println(databaseURL)
	if *conn_str == "" {
		fmt.Println("Missing -cs (connection string) parameter")
		usage()
		os.Exit(1)
	}
	dbconn, err := sql.Open("oracle", *conn_str)
	if err != nil {
		fmt.Println("Can't open the connection ", err)
		return
	}
	defer func() {
		err = dbconn.Close()
		if err != nil {
			fmt.Println("Can't close connection: ", err)
		}
	}()

	err = dbconn.Ping()
	if err != nil {
		fmt.Println("Can't ping the database ", err)
	}

}
