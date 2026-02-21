package hash_table

import (
	"math"
)

type HashFuncType int
type HashFunc func(int, int) int

const (
	DIVISION HashFuncType = iota
	MULTIPLICATION
)

func SelectHashFunc(t HashFuncType) HashFunc {
	switch t {
	case MULTIPLICATION:
		return func(k, m int) int { return multiplicationHash(k, m, 0.5) }
	default:
		return func(k, m int) int { return divisonHash(k, m) }
	}
}

func divisonHash(k, m int) int {
	return k % m
}

func multiplicationHash(k, m int, A float64) int {
	return int(math.Floor(float64(m) * ((float64(k) * A) - math.Floor(float64(k)*A))))
}
