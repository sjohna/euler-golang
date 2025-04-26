package euler51

import (
	"euler/utility/prime"
	"strconv"
	"strings"
)

func Iterate() int {
	type Entry struct {
		count    int
		smallest int
	}

	templates := make(map[string]*Entry)

	primes := prime.NaiveGenerator().Infinite()

	for {
		currPrime := primes()
		templatesForPrime := ToTemplates(currPrime)
		for _, t := range templatesForPrime {
			if e, ok := templates[t]; ok {
				e.count += 1
				if e.count == 8 {
					return e.smallest
				}
			} else {
				templates[t] = &Entry{
					count:    1,
					smallest: currPrime,
				}
			}
		}
	}
}

func ToTemplates(n int) []string {
	nStr := strconv.Itoa(n)

	var subs []string
	for c := uint8('0'); c <= uint8('9'); c++ {
		var subsForThis []string
		var anyReplaced bool
		for i := range nStr {
			if nStr[i] == c {
				sub := []byte(strings.Clone(nStr))
				sub[i] = '_'
				anyReplaced = true

				// clone and replace this in each of the already-found subs
				var newSubs []string
				for _, s := range subsForThis {
					newS := []byte(strings.Clone(s))
					newS[i] = '_'
					newSubs = append(newSubs, string(newS))
				}
				subsForThis = append(subsForThis, newSubs...)
				subsForThis = append(subsForThis, string(sub))
			}
		}

		if anyReplaced {
			subs = append(subs, subsForThis...)
		}
	}

	return subs
}
