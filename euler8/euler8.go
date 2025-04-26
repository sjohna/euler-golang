package euler8

import (
	"euler/utility"
	"euler/utility/sequence"
	"euler/utility/slice"
	"io"
	"os"
)

/*
https://projecteuler.net/problem=8

What is the largest product of 13 consecutive digits in the 1000-digit number in number.txt?
*/

func LoadArray() ([]int, error) {
	file, err := os.Open("number.txt")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var result []int
	buf := make([]byte, 1024)

	for read, err := file.Read(buf); read > 0 && (err == nil || err == io.EOF); read, err = file.Read(buf) {
		for i := 0; i < read; i++ {
			if '0' <= buf[i] && buf[i] <= '9' {
				result = append(result, int(buf[i]-'0'))
			}
		}
	}

	return result, nil
}

func Naive() int {
	arr, _ := LoadArray()

	maxProduct := 0
	for i := 0; i < len(arr); i++ {
		product := 1
		for j := i; j < len(arr) && j < i+13; j++ {
			product *= arr[j]
		}
		if product > maxProduct {
			maxProduct = product
		}
	}

	return maxProduct
}

func Generator() int {
	arr, _ := LoadArray()

	return sequence.Range(0, len(arr)-13).Map(func(n int) int {
		return slice.Product(arr[n : n+13])
	}).Reduce(utility.Max, 0)
}
