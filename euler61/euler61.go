package euler61

import (
	u "euler/utility"
	"fmt"
	"slices"
)

func Adjacent(a, b int) bool {
	return a%100 == b/100
}

func SolveIt() int {
	triangles := u.NaturalNumbers().Map(u.Triangular).SkipWhile(u.LessThan(1000)).TakeWhile(u.LessThan(10000)).ToSlice()
	squares := u.NaturalNumbers().Map(u.Square[int]).SkipWhile(u.LessThan(1000)).TakeWhile(u.LessThan(10000)).ToSlice()
	pentagons := u.NaturalNumbers().Map(u.Pentagonal).SkipWhile(u.LessThan(1000)).TakeWhile(u.LessThan(10000)).ToSlice()
	hexagons := u.NaturalNumbers().Map(u.Hexagonal).SkipWhile(u.LessThan(1000)).TakeWhile(u.LessThan(10000)).ToSlice()
	heptagons := u.NaturalNumbers().Map(u.Heptagonal).SkipWhile(u.LessThan(1000)).TakeWhile(u.LessThan(10000)).ToSlice()
	octagons := u.NaturalNumbers().Map(u.Octagonal).SkipWhile(u.LessThan(1000)).TakeWhile(u.LessThan(10000)).ToSlice()

	//fmt.Println(len(triangles), len(squares), len(pentagons), len(hexagons), len(heptagons), len(octagons))

	// cull those with no connections
	in := make(map[int]bool)
	out := make(map[int]bool)

	numLists := [][]int{triangles, squares, pentagons, hexagons, heptagons, octagons}

	for i := range numLists {
		for j := range numLists {
			if i == j {
				continue
			}
			for _, v1 := range numLists[i] {
				for _, v2 := range numLists[j] {
					if Adjacent(v1, v2) {
						out[v1] = true
						in[v2] = true
					}
				}
			}
		}
	}

	connectedTriangles := CullUnconnected(triangles, in, out)
	connectedSquares := CullUnconnected(squares, in, out)
	connectedPentagons := CullUnconnected(pentagons, in, out)
	connectedHexagons := CullUnconnected(hexagons, in, out)
	connectedHeptagons := CullUnconnected(heptagons, in, out)
	connectedOctagons := CullUnconnected(octagons, in, out)

	//fmt.Println(len(connectedTriangles), len(connectedSquares), len(connectedPentagons), len(connectedHexagons), len(connectedHeptagons), len(connectedOctagons))

	numLists = [][]int{connectedSquares, connectedPentagons, connectedHexagons, connectedHeptagons, connectedOctagons}

	for _, value := range connectedTriangles {
		if sum := Recurse(numLists, []int{value}); sum != -1 {
			return sum
		}
	}

	return -1
}

func Recurse(numLists [][]int, nums []int) int {
	lastNum := nums[len(nums)-1]
	if len(numLists) == 0 {
		if Adjacent(lastNum, nums[0]) {
			fmt.Println(nums)
			return u.SliceGenerator(nums).Reduce(u.Sum, 0)
		} else {
			return -1
		}
	}

	newNums := append(slices.Clone(nums), 0)

	for i := range numLists {
		for j := range numLists[i] {
			if Adjacent(lastNum, numLists[i][j]) {
				newLists := append(slices.Clone(numLists[:i]), numLists[i+1:]...)
				newNums[len(newNums)-1] = numLists[i][j]
				sum := Recurse(newLists, newNums)
				if sum != -1 {
					return sum
				}
			}
		}
	}

	return -1
}

func CullUnconnected(nums []int, in, out map[int]bool) []int {
	ret := make([]int, 0)
	for _, n := range nums {
		if in[n] && out[n] {
			ret = append(ret, n)
		}
	}

	return ret
}
