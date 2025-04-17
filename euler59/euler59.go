package euler59

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadCipherText() []byte {
	raw, _ := os.ReadFile("euler59/0059_cipher.txt")
	rawStr := string(raw)
	tokens := strings.Split(rawStr, ",")

	var chars []byte
	for _, token := range tokens {
		char, _ := strconv.Atoi(token)
		chars = append(chars, uint8(char))
	}

	return chars
}

// 20 most common English words, from https://en.wikipedia.org/wiki/Most_common_words_in_English
var commonWords = []string{"the", "be", "to", "of", "and", "a", "in", "that", "have", "i", "it", "for", "not", "on", "with", "he", "as", "you", "do", "at"}

func CommonWordsMap() map[string]bool {
	ret := make(map[string]bool)
	for _, word := range commonWords {
		ret[word] = true
	}

	return ret
}

func SolveIt() int {
	cipherText := LoadCipherText()
	commonWords := CommonWordsMap()

	mostCommonWords := 0
	var bestDecipher string

	for c1 := 'a'; c1 <= 'z'; c1++ {
		for c2 := 'a'; c2 <= 'z'; c2++ {
			for c3 := 'a'; c3 <= 'z'; c3++ {
				key := []byte{byte(c1), byte(c2), byte(c3)}

				deciphered := string(XorCipher(cipherText, key))

				commonWordCount := 0
				for _, word := range strings.Split(deciphered, " ") {
					if commonWords[strings.ToLower(word)] {
						commonWordCount++
					}
				}

				if commonWordCount > mostCommonWords {
					mostCommonWords = commonWordCount
					bestDecipher = deciphered
					fmt.Printf("%d common words (%s): %s\n", commonWordCount, string(key), deciphered)
				}
			}
		}
	}

	sum := 0
	for _, char := range []byte(bestDecipher) {
		sum += int(char)
	}

	return sum
}

func XorCipher(text, key []byte) []byte {
	ret := make([]byte, len(text))
	for i := range text {
		ret[i] = text[i] ^ key[i%len(key)]
	}

	return ret
}
