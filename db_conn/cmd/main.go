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
package main

import (
	"db_conn"
	"flag"
	"log"
)

var (
	dbPwd    = flag.String("pw", "", "Database user password")
	dbUser   = flag.String("u", "", "Database user")
	dbHost   = flag.String("h", "", "Database host")
	dbPort   = flag.Int("p", 1521, "Database instance listener port - 1521 default for Oracle")
	dbName   = flag.String("s", "", "Database name or service(for Oracle)")
	dbWallet = flag.String("w", "", "Path to unzipped Oracle wallet (for SSL connection)")
)

func main() {
	flag.Parse()
	if *dbUser == "" || *dbPwd == "" || *dbPort <= 0 || *dbHost == "" || *dbName == "" {
		db_conn.Usage()
		log.Fatal("Missing one of the required options for connection", *dbHost, *dbPort, *dbName, *dbUser, *dbUser, "*******")
	}
	db_conn.RunApp(*dbUser, *dbPwd, *dbHost, *dbPort, *dbName, *dbWallet)
}
