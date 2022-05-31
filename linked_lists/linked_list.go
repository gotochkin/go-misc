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

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

func (l linkedList) Len() int {
	return l.length
}

func (l linkedList) Display() {
	for l.head != nil {
		fmt.Printf("%v -> ", l.head.data)
		l.head = l.head.next
	}
	fmt.Println()
}

func (l *linkedList) PushBack(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.next = n
		l.tail = n
		l.length++
	}
}

func (l *linkedList) Delete(key int) {

	if l.head.data == key {
		l.head = l.head.next
		l.length--
		return
	}
	var prev *node = nil
	curr := l.head
	for curr != nil && curr.data != key {
		prev = curr
		curr = curr.next
	}
	if curr == nil {
		fmt.Println("Key Not found")
		return
	}
	prev.next = curr.next
	l.length--
	fmt.Println("Node Deleted")
}
func (l *linkedList) Front() (int, error) {
	//
	if l.head == nil {
		return 0, fmt.Errorf("No Front value in an empty list")
	}
	return l.head.data, nil
}

func (l *linkedList) Back() (int, error) {
	if l.head == nil {
		return 0, fmt.Errorf("No Back value in an empty list")
	}
	return l.tail.data, nil
}

var (
	cv         = flag.Int("v", 1, "Node value to search/delete")
	inputArray arrayVars
)

func main() {
	//
	flag.Var(&inputArray, "ia", "Input comma separated array of values")
	flag.Parse()
	mylist := linkedList{}
	for _, i := range inputArray {
		k, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		node := &node{data: k}
		mylist.PushBack(node)
	}

	mylist.Display()
	mylist.Delete(*cv)
	mylist.Display()
}
