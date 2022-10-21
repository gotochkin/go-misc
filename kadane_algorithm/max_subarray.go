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
/// Search subarray with max sum of members in an array
package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type arrayVars []string

func (i *arrayVars) String() string {
	return "String of parameters"
}

func (i *arrayVars) Set(s string) error {
	*i = strings.Split(s, ",")
	return nil
}

var (
	inputArray arrayVars
)

func findMaxArray(ia []int) (int, int, int, error) {
	// Search
	curr_max := ia[0]
	global_max := ia[0]
	s := 0
	e := 0
	for i := range ia {
		if curr_max < 0 {
			curr_max = ia[i]
			s = i
		} else {
			curr_max += ia[i]
		}
		if global_max < curr_max {
			global_max = curr_max
			e = i
		}
		//fmt.Printf("Index for the first and last members in subarray %d %d current_max %d global max %d\n", s, e, curr_max, global_max)
	}
	return s, e, global_max, nil
}

func main() {
	//
	fmt.Println("Starting")
	// Convert string array to integer
	flag.Var(&inputArray, "ia", "Input comma separated array of values")
	flag.Parse()
	var ia = []int{}
	for _, i := range inputArray {
		k, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		ia = append(ia, k)
	}
	s, e, m, err := findMaxArray(ia)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arrray starts on element index %d, finishes on %d and the sum of members is %d \n", s, e, m)
}
