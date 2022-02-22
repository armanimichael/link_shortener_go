package shortener

import "math/rand"

func getAsciiChars(index int, count int) []byte {
	var list []byte

	current := index
	for i := 0; i < count; i++ {
		list = append(list, byte(current))
		current++
	}

	return list
}

func getAllowedChars() []byte {
	const lettersCount = 26
	uppercase := getAsciiChars('A', lettersCount)
	lowercase := getAsciiChars('a', lettersCount)
	numbers := getAsciiChars('0', 10)

	letters := append(lowercase, uppercase...)
	return append(numbers, letters...)
}

// GetRandomCombination generates a random combination of bytes of a given length
func GetRandomCombination(length int) []byte {
	var combination []byte
	allowed := getAllowedChars()
	allowedLen := len(allowed)

	for i := 0; i < length; i++ {
		random := rand.Intn(allowedLen)
		combination = append(combination, allowed[random])
	}

	return combination
}
