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
/// Package contains code to multiply files based on a template
/// The example is created based on an ArgoCD application template

package main

import (
	"flag"
	"fmt"
	"sort"
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
func double_point(iar []int, v int) (int, error) {
	//
	sort.Ints(iar)
	l := 0
	r := len(iar) - 1
	for l < r {
		//fmt.Println(iar[l] + iar[r])
		if iar[l]+iar[r] == v {
			fmt.Println("sum of " + strconv.Itoa(iar[l]) + " and " + strconv.Itoa(iar[r]))
			return 1, nil
		} else if iar[l]+iar[r] < v {
			l++
		} else {
			r--
		}
	}
	return 0, nil
}

func map_search(iar []int, v int) (int, error) {
	//
	v_map := make(map[int]int)
	for _, ia := range iar {
		//
		//fmt.Println(v_map)
		if val, ok := v_map[v-ia]; ok {
			fmt.Println("sum of " + strconv.Itoa(val) + " and " + strconv.Itoa(ia))
			return 1, nil
		}
		v_map[ia] = ia

	}
	return 0, nil
}

var (
	cv         = flag.Int("v", 1, "Value to compare")
	inputArray arrayVars
)

func main() {
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
	// Uncomment the double_points and comment the map_search if you want to change the function
	//k, err := double_point(ia, *cv)
	k, err := map_search(ia, *cv)
	if err != nil {
		panic(err)
	}

	//fmt.Println(ia)
	fmt.Println(k)

}
