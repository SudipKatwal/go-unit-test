package randomnumber

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(min, max int) int {
	if max < min {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
