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
	return random.random.Intn(max)
}

func (random *Random) RandomElement(elements []interface{}) interface{} {
	max := len(elements)
	randomNumber := random.random.Intn(max)

	return elements[randomNumber]
}
