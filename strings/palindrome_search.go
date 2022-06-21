package main

import "fmt"

func palindrome_search(ia string) {
	//
	for i := 0; i < len(ia); i++ {
		if i == 0 {
			//fmt.Println(ia[0 : i+1])
			palindrome_search_substrings(ia, 0, i+1)
		} else {
			fmt.Println(ia[i : i+1])
			palindrome_search_substrings(ia, i, i+1)
			//fmt.Println(ia[i-1 : i+1])
			palindrome_search_substrings(ia, i-1, i+1)
		}
	}

}
func palindrome_search_substrings(ia string, j int, k int) {
	//
	for j >= 0 && k < len(ia) {
		if ia[j] != ia[k] {
			break
		}
		fmt.Println(ia[j : k+1])
		j--
		k++

	}
}
func main() {
	in_array := [...]string{"abracadabra", "abba", "madam"}
	for _, g := range in_array {
		palindrome_search(g)
	}
}
