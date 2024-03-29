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
//// Search a substring existing in a dictionary
package main

import "fmt"

func main() {
	dic_array := [...]string{"monkey", "test", "hello", "fish", "apple"}
	input_string := "wgtrapplejfishkmonkey"
	for i := 0; i <= len(input_string); i++ {
		for j := i; j <= len(input_string); j++ {
			for m := range dic_array {
				if input_string[i:j] == dic_array[m] {
					fmt.Println(input_string[i:j])
				}
			}
		}
	}
}
