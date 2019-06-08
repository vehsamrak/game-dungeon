package random

import (
	"math/rand"
	"time"
)

type Random struct {
	random *rand.Rand
}

func (random Random) Create() *Random {
	return &Random{random: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (random *Random) Seed(seed int64) {
	random.random.Seed(seed)
}

func (random *Random) RandomNumber(max int) int {
	randomNumber := random.random.Intn(max + 1)

	if randomNumber > max {
		return random.RandomNumber(max)
	}

	return randomNumber
}

func (random *Random) RandomBoolean() bool {
	return random.RandomNumber(1) == 1
}
