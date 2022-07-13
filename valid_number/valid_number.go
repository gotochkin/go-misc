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
/// Example of parsing different arguments including an array of values

package main

import (
	"flag"
	"fmt"
)

var (
	inputString = flag.String("ia", "1.25", "Put a string to check if the string is valid number")
)

func main() {
	flag.Parse()
	st := ""
	if len(*inputString) > 0 {
		for i := range *inputString {
			if *inputString[0:1] == "+" || *inputString[0:1] == "-" {
				if len(*inputString) < 2 {
					fmt.Println("Not valid number")
				}
				st := *inputString[1:2]
			}
			// if i[0] == "-" || i[0] == "+" {
			// 	st = i[1]
			// } else {
			// 	st = i[0]
			// }
			//st = i[0]
			fmt.Printf("%i", str)

		}
		fmt.Println(st)
	}

}
