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
/// Binary tree implementation with comaprison of two binary trees

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
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
	//cv         = flag.Int("v", 1, "Node value to search/delete")
	inputArray arrayVars
)

type Bnode struct {
	left  *Bnode
	right *Bnode
	data  int64
}

type Btree struct {
	root *Bnode
}

func (t *Btree) insert(data int64) *Btree {
	if t.root == nil {
		t.root = &Bnode{
			data:  data,
			left:  nil,
			right: nil,
		}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *Bnode) insert(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &Bnode{
				data:  data,
				left:  nil,
				right: nil,
			}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Bnode{
				data:  data,
				left:  nil,
				right: nil,
			}
		} else {
			n.right.insert(data)
		}
	}
}

func (t *Btree) mirrorTree(n *Bnode) {
	if n != nil {
		t.mirrorTree(n.left)
		t.mirrorTree(n.right)
		var ntmp *Bnode = n.left
		n.left = n.right
		n.right = ntmp
	}
}

func printtree(w io.Writer, node *Bnode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	printtree(w, node.left, ns+2, 'L')
	printtree(w, node.right, ns+2, 'R')
}

func main() {
	flag.Var(&inputArray, "ia", "Input comma separated array of values")
	flag.Parse()
	tree := &Btree{}
	for _, i := range inputArray {
		//k, err := strconv.Atoi(i)
		k, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			panic(err)
		}
		tree.insert(k)
	}
	printtree(os.Stdout, tree.root, 0, 'M')
	tree.mirrorTree(tree.root)
	printtree(os.Stdout, tree.root, 0, 'M')
}
