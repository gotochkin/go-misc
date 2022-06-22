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
// Search all palindromes in an array of strings

package main

import "fmt"

func palindrome_search(ia string) {
	//
	for i := 0; i < len(ia); i++ {
		if i == 0 {
			//fmt.Println(ia[0 : i+1])
			palindrome_search_substrings(ia, 0, i+1)
		} else {
			fmt.Println(ia[i : i+1])
			palindrome_search_substrings(ia, i, i+1)
			//fmt.Println(ia[i-1 : i+1])
			palindrome_search_substrings(ia, i-1, i+1)
		}
	}

}
func palindrome_search_substrings(ia string, j int, k int) {
	//
	for j >= 0 && k < len(ia) {
		if ia[j] != ia[k] {
			break
		}
		fmt.Println(ia[j : k+1])
		j--
		k++

	}
}
func main() {
	in_array := [...]string{"abracadabra", "abba", "madam"}
	for _, g := range in_array {
		palindrome_search(g)
	}
}
