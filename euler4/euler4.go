package main

import (
	"fmt"
	"sort"
	"strconv"
)

func IsPalindrome(num int) bool {
	numstr := strconv.Itoa(num)
	runes := []rune(numstr) // can I access the characters directly?

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}

func Euler4() int {
	products := make([]int, 0)

	for n1 := 1; n1 < 1000; n1++ {
		for n2 := n1; n2 < 1000; n2++ {
			products = append(products, n1*n2)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(products)))

	for i := 0; i < len(products); i++ {
		if IsPalindrome(products[i]) {
			return products[i]
		}
	}

	return -1
}

func main() {
	fmt.Println(Euler4())
}
