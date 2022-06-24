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
/// Max number of ids in the intersection of ranges
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz1234567890"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RndString(length int) string {
	return StringWithCharset(length, charset)
}

func main() {
	//Initial data
	//Array
	arr := [30][2]int{}
	//Maps
	ven_map := map[string]int{}
	vex_map := map[string]int{}
	//id
	var id_str string
	// Filling the arrays and maps
	for i := 0; i < 11; i++ {
		arr[i][0] = i
		arr[i][1] = arr[i][0] + rand.Intn(3)
		id_str = RndString(5)
		ven_map[id_str] = i
		vex_map[id_str] = i + rand.Intn(3)
		fmt.Println(id_str + "	" + strconv.Itoa(ven_map[id_str]) + "	" + strconv.Itoa(vex_map[id_str]))
	}
	vst_map := map[int]int{}
	for key, k := range ven_map {
		t := vex_map[key] - k
		for j := k; j < k+t+1; j++ {
			vst_map[j]++
		}
		//fmt.Println(key + "	" + strconv.Itoa(t))
	}
	l := 0
	m := 0
	for key, v := range vst_map {
		//fmt.Println(strconv.Itoa(key) + "	" + strconv.Itoa(v))
		if v > l {
			l = v
			m = key
		}
	}
	fmt.Println("Max ids in the key: " + strconv.Itoa(m))

}
