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
/// Move all zeroes to left in the array keeping the order

package main

import (
	"flag"
	"fmt"
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

func main() {
	flag.Var(&inputArray, "ia", "Input comma separated array of values")
	flag.Parse()
	writeidx := len(inputArray) - 1
	for readidx := len(inputArray) - 1; readidx >= 0; readidx-- {
		//fmt.Println(inputArray[readidx])
		if inputArray[readidx] != "0" {
			inputArray[writeidx] = inputArray[readidx]
			fmt.Println(inputArray[readidx])
			writeidx--
		}
	}
	fmt.Println(writeidx)
	for i := 0; i <= writeidx; i++ {
		inputArray[i] = "0"
	}
	fmt.Println(inputArray)

}
