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
// Test for comapre two binary trees

package main

import (
	"testing"
)

func TestCompareTree(t *testing.T) {
	var tests = []struct {
		tarray1 []int64
		tarray2 []int64
		result  bool
	}{
		{[]int64{-1, -5, 0, 32, -2, 3, -4}, []int64{-1, -5, 0, 32, -2, 3, -4}, true},
		{[]int64{-2, 4, 0, -1, -2, 11, -8}, []int64{-2, 4, 0, -1, -2, 11, -8}, true},
		{[]int64{1, 2, -2, 0, -2, 1, -3, 5, -6}, []int64{1, 2, -2, 0, -2, 1, 3, 5, -6}, true},
	}
	for _, test := range tests {
		tree1 := &Btree{}
		tree2 := &Btree{}
		for _, k := range test.tarray1 {
			tree1.insert(k)
		}
		for _, k := range test.tarray2 {
			tree2.insert(k)
		}
		if res := compareTree(tree1.root, tree2.root); res != test.result {
			t.Errorf("For compareTree(%v,%v) expected %v but got %v", test.tarray1, test.tarray2, test.result, res)
		}
	}
}
