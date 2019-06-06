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

func (random *Random) RandomNumber(max int64) int64 {
	return random.random.Int63n(max)
}
