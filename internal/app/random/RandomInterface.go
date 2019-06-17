package random

type Random interface {
	Seed(seed int64)
	RandomNumber(max int) int
	RandomBoolean() bool
}
