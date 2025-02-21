package euler22

import (
	"os"
	"sort"
	"strings"
)

/*
https://projecteuler.net/problem=22

We calculate the score of a name as the sum of the alphabetical values of its letters, multiplied by the name's position
in the list of all names sorted alphabetically.

What is the total score of all the names in names.txt?
*/

func GetNamesList() []string {
	fileContent, _ := os.ReadFile("names.txt")

	fileContentString := string(fileContent)

	tokens := strings.Split(fileContentString, ",")

	var names []string

	for _, rawName := range tokens {
		name := strings.Trim(rawName, "\" ,")
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

func AlphabetScore(s string) int {
	bytes := []byte(s)
	score := 0
	for _, b := range bytes {
		score += int((b - 'A') + 1)
	}

	return score
}

func Loop() int {
	names := GetNamesList()

	totalScore := 0
	for i, name := range names {
		totalScore += AlphabetScore(name) * (i + 1)
	}

	return totalScore
}
