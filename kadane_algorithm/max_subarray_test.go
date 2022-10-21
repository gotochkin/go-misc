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
/// Test for max_subarray.go
package main

import (
	"testing"
)

//!+test

func TestFindMaxArray(t *testing.T) {
	var tests = []struct {
		tarray   []int
		startInd int
		endInd   int
		sumVal   int
	}{
		{[]int{-1, -5, 0, 32, -2, 3, -4}, 2, 5, 33},
		{[]int{-2, 4, 0, -1, -2, 11, -8}, 1, 5, 12},
		{[]int{1, 2, -2, 0, -2, 1, -3, 5, -6}, 7, 7, 5},
	}
	for _, test := range tests {
		s, e, sm, err := findMaxArray(test.tarray)
		if err != nil {
			t.Errorf("findMaxArray failed: %v", err)
		} else if s != test.startInd || e != test.endInd || sm != test.sumVal {
			t.Errorf("find_max_array(%q) returned %d %d %d when expected %d %d %d", test.tarray, s, e, sm, test.startInd, test.endInd, test.sumVal)
		}
	}
}

//!-test
