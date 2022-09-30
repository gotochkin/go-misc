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
package db_conn

import (
	"database/sql"
	"fmt"

	go_ora "github.com/sijms/go-ora/v2"
)

func connectOracle(dbUser string, dbPwd string, dbHost string, dbPort int, dbName string, dbWallet string) (*sql.DB, error) {
	// Connection parameters
	// Here is example how to use environment variable for that, but it is not secure
	// Better to use integrations with secret stores
	var (
		err         error
		databaseURL string
	)
	//create URL for the database connection

	if dbWallet == "" {
		databaseURL = go_ora.BuildUrl(dbHost, dbPort, dbName, dbUser, dbPwd, nil)
	} else {
		//
		databaseURLopt := make(map[string]string)
		databaseURLopt["TRACE FILE"] = "ora_conn.log" //can be another parameter
		databaseURLopt["SSL"] = "enable"
		databaseURLopt["SSL Verify"] = "false"
		databaseURLopt["WALLET"] = dbWallet
		databaseURL = go_ora.BuildUrl(dbHost, dbPort, dbName, dbUser, dbPwd, databaseURLopt)

	}

	//Create a connection pool
	dbPool, err := sql.Open("oracle", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("sql open pool error: %v", err)
	}

	//Configure the pool
	configurePool(dbPool)

	return dbPool, nil

}
