package shortener

import "testing"

const length int = 10
const allowedCharsCount int = 62

func Test_AllowedChars(t *testing.T) {
	allowed := getAllowedChars()

	if len(allowed) != allowedCharsCount {
		t.Errorf("Allowed characters should be %d", allowedCharsCount)
	}
}

func Test_RandomCombination(t *testing.T) {
	combination := GetRandomCombination(length)

	combinationToStr := ""
	for _, item := range combination {
		combinationToStr += string(item)
	}
	t.Logf("\n%s", combinationToStr)
}

func Test_Randomicity(t *testing.T) {
	var combinations [1000000]string

	for i := 0; i < len(combinations); i++ {
		random := GetRandomCombination(10)
		combinations[i] = string(random)
	}

	collisionCount := 0
	collisions := map[string]int{
		combinations[0]: 0,
	}

	for i, item := range combinations {
		if i == 0 {
			continue
		}

		if _, exists := collisions[item]; !exists {
			collisions[item] = i
		} else {
			collisionCount++
		}
	}

	t.Logf("\nCollisions: %d", collisionCount)
}
