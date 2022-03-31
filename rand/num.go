package rand

import (
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

type Num interface {
	constraints.Float | constraints.Integer
}

// Rand return a random num between [low, up),
// support all num type,
// if low > up, it will return 0.
func RandNum[T Num](low, up T) T {
	if low > up {
		return 0
	}
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	return T(randomizer.Int63n(int64(up-low)) + int64(low))
}
