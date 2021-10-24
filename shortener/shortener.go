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
	uppercase := getAsciiChars(97, 26)
	lowercase := getAsciiChars(65, 26)
	numbers := getAsciiChars(48, 10)

	letters := append(lowercase, uppercase...)
	return append(numbers, letters...)
}

func GetRandomCombination(length int) []byte {
	var list []byte
	allowed := getAllowedChars()
	allowedLen := len(allowed)

	for i := 0; i < length; i++ {
		random := rand.Intn(allowedLen)
		list = append(list, allowed[random])
	}

	return list
}
