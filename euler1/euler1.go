package euler1

/*
https://projecteuler.net/problem=1

Find the sum of all the multiples of 3 or 5 below 1000.
*/

func Loop() int {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
