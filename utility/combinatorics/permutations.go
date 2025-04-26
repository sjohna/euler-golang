package combinatorics

import "fmt"

func Heap[T any](k int, A []T) {
	if k == 1 {
		fmt.Println(A)
	} else {
		Heap[T](k-1, A)

		for i := range k - 1 {
			if k%2 == 0 {
				A[i], A[k-1] = A[k-1], A[i]
			} else {
				A[0], A[k-1] = A[k-1], A[0]
			}
			Heap[T](k-1, A)
		}
	}
}

func KthPermutationOfNItems(N, K int) []int {
	K = K - 1

	coefficients := make([]int, N)

	for i := range N {
		coefficients[N-i-1] = K % (i + 1)
		K /= i + 1
	}

	for i := N - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if coefficients[j] <= coefficients[i] {
				coefficients[i]++
			}
		}
	}

	return coefficients
}
