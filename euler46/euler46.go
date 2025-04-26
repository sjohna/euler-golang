package euler46

import (
	"container/heap"
	"euler/utility/prime"
)

/*
A less-famous Goldbach conjecture is that every odd composite number can be written as a prime plus twice a square.

What is the smallest counterexample to this conjecture?
*/

type PrimeSquareSum struct {
	prime  int
	square int
}

func (p PrimeSquareSum) Value() int {
	return p.prime + 2*p.square*p.square
}

type CandidateList struct {
	list []PrimeSquareSum
}

// implement sort.Interface and heap.Interface for CandidateList

func (l *CandidateList) Len() int {
	return len(l.list)
}

func (l *CandidateList) Less(i, j int) bool {
	return l.list[i].Value() < l.list[j].Value()
}

func (l *CandidateList) Swap(i, j int) {
	l.list[i], l.list[j] = l.list[j], l.list[i]
}

func (l *CandidateList) Push(v any) {
	l.list = append(l.list, v.(PrimeSquareSum))
}

func (l *CandidateList) Pop() any {
	ret := l.list[len(l.list)-1]
	l.list = l.list[:len(l.list)-1]
	return ret
}

func Heap() int {
	oddPrimes := prime.CachedGenerator().Skip(1).Infinite()

	frontier := &CandidateList{}

	// push 3 onto Frontier to initialize
	heap.Push(frontier, PrimeSquareSum{
		oddPrimes(),
		0,
	})

	maxSeen := 1

	for {
		nextCandidate := heap.Pop(frontier).(PrimeSquareSum)
		nextValue := nextCandidate.Value()

		if nextValue > maxSeen+2 {
			// found a gap
			return maxSeen + 2
		}

		if nextValue > maxSeen {
			maxSeen = nextValue
		}

		heap.Push(frontier, PrimeSquareSum{
			nextCandidate.prime,
			nextCandidate.square + 1,
		})

		if nextCandidate.square == 0 {
			heap.Push(frontier, PrimeSquareSum{
				oddPrimes(),
				0,
			})
		}
	}
}
