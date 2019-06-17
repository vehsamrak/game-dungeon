package random

import (
	"math/rand"
	"time"
)

type Randomizer struct {
	random *rand.Rand
}

func (random Randomizer) Create() *Randomizer {
	return &Randomizer{random: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (random *Randomizer) Seed(seed int64) {
	random.random.Seed(seed)
}

func (random *Randomizer) RandomNumber(max int) int {
	randomNumber := random.random.Intn(max + 1)

	if randomNumber > max {
		return random.RandomNumber(max)
	}

	return randomNumber
}

func (random *Randomizer) RandomBoolean() bool {
	return random.RandomNumber(1) == 1
}
