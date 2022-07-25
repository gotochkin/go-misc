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
/// Binary search in a sorted array
/// Values in the array are integer and can be negative numbers
//
package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// Type for input array of integers as ann argument
type arrayVars []string

// Methods String and Set for the type arrayVars
func (i *arrayVars) String() string {
	return "Array of string"
}

func (i *arrayVars) Set(s string) error {
	*i = strings.Split(s, ",")
	return nil
}

//Binary search in the array
func b_search(ia []int, sv int) (int, error) {
	// Initializing
	lv := 0
	hv := len(ia)

	if hv < lv || ia[lv] > sv || ia[hv-1] < sv {
		return -1, nil
	}
	for hv > lv {
		//Get the middle
		if ia[lv] == sv {
			return lv, nil
		}
		mid := lv + (hv-lv)/2

		if ia[mid] < sv {
			lv = mid + 1
		} else if ia[mid] > sv {
			hv = mid
		} else if ia[mid] == sv {
			return mid, nil
		}
	}
	return -2, nil
}

// Variables
var (
	sv         = flag.Int("sv", 1, "Search value")
	inputArray arrayVars
)

func main() {
	//Parse variables
	flag.Var(&inputArray, "ia", "Input comma separated list of sorted integer values")
	flag.Parse()

	//Convert array of strings to integers
	var ia = []int{}
	for _, i := range inputArray {
		k, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		ia = append(ia, k)
	}
	// Call a function to search the value in the array
	k, err := b_search(ia, *sv)
	if err != nil {
		panic(err)
	}
	fmt.Println(k)

}
