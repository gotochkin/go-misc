package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"flag"
	"io"

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
	inputArray1 arrayVars
	inputArray2 arrayVars
)

type Bnode struct {
	left *Bnode
	right *Bnode
	data int64
}

type Btree struct {
	root *Bnode
}

func (t *Btree) insert (data int64) *Btree {
	if t.root == nil {
		t.root = &Bnode {
			data: data,
			left: nil,
			right: nil,
		}
	}else{
		t.root.insert(data)
	}
	return t
}

func (n *Bnode) insert (data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &Bnode {
				data: data,
				left: nil,
				right: nil,
			}
		}else{
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Bnode {
				data: data,
				left: nil,
				right: nil,
			}
		} else {
			n.right.insert(data)
		}
	}	
}

func(t *Btree) mirrorTree(n *Bnode) {
	if n != nil {
		t.mirrorTree(n.left)
		t.mirrorTree(n.right)
		var ntmp *Bnode = n.left
		n.left = n.right
		n.right = ntmp
	}
}

func compareTree(n1 *Bnode, n2 *Bnode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 != nil && n1 != nil {
		return (( n1.data == n2.data ) && 
		compareTree(n1.left,n2.left) && compareTree(n1.right,n2.right))
	}
	return false
}

func printtree(w io.Writer, node *Bnode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	printtree(w,node.left,ns+2,'L')
	printtree(w, node.right,ns+2,'R')
}

func main() {
	flag.Var(&inputArray1, "ia1", "Input comma separated array of values")
	flag.Var(&inputArray2, "ia2", "Input comma separated array of values")
	flag.Parse()
	tree1 := &Btree{}
	tree2 := &Btree{}
	for _, i := range inputArray1 {
		//k, err := strconv.Atoi(i)
		k, err := strconv.ParseInt(i,10,64)
		if err != nil {
			panic(err)
		}
		tree1.insert(k)
	}
	for _, i := range inputArray2 {
		//k, err := strconv.Atoi(i)
		k, err := strconv.ParseInt(i,10,64)
		if err != nil {
			panic(err)
		}
		tree2.insert(k)
	}
	printtree(os.Stdout,tree1.root,0,'M')
	//tree.mirrorTree(tree.root)
	if compareTree(tree1.root,tree2.root) {
		fmt.Println("The trees are identical")
	} else {
		fmt.Println("The trees are not identical")
	}
	printtree(os.Stdout,tree2.root,0,'M')
}