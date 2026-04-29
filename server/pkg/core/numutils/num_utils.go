// Package numutils ...
package numutils

import (
	"math/rand"
	"time"
)

// RandInt ...
func RandInt(minNum int, maxNum int) int {
	return minNum + rand.Intn(maxNum-minNum)
}

// RandDelay ...
func RandDelay(minNum int, maxNum int) time.Duration {
	return time.Duration(minNum + rand.Intn(maxNum-minNum))
}

// GetPercentValue ...
func GetPercentValue(number int, percent int) int {
	return (number / 100) * percent
}
